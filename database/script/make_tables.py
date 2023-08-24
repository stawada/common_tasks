import os
import pandas as pd
import random
from hashlib import sha256
from googletrans import Translator
import time
from tqdm import tqdm

def make_student_table(student_df):
    pairs = []
    pairs_txt = "" # 学生IDとパスワード(平文)を保持
    res = [] # テーブルに必要なデータを保持
    for i in range(student_df.shape[0]):
        up_char = random.choice([chr(j).upper() for j in range(97,110)])
        student_id = up_char + str(i+1).zfill(9)
        student_name = student_df.loc[i, "student_name"]
        pw = "".join(random.choices([chr(j) for j in range(33,127)], k=10))
        hashed_pw = sha256(pw.encode()).hexdigest() # ハッシュ化
        pairs.append([student_id, pw])
        pairs_txt += f"{student_id}, {pw}\n"
        res.append([student_id, student_name, hashed_pw])

    return pairs, pairs_txt, res

def make_lect_ctlg_table(subject_df, subject_n):
    res = []
    translator = Translator()
    for i in range(subject_df.shape[0]):
        subject_name = subject_df.loc[i, "lecture_name"]
        en_subject_name = translator.translate(subject_name, src="ja", dest="en").text # ja -> en
        for j in range(subject_n):
            pascal_str = ""
            for word in en_subject_name.split(" "):
                if word in ["Ⅰ", "Ⅱ", "Ⅲ"]:
                    pascal_str += word
                else:
                    pascal_str += word[0].upper() + word[1:]
            lect_ctlg_id = f"{pascal_str}-{str(j+1).zfill(2)}"
            res.append([lect_ctlg_id, subject_name])

    return res

def make_lect_hist_table(subject_df, subject_n, ctlg_elms):
    now_unix_time = int(time.time())
    res = []
    for i in range(len(subject_df)):
        for j in range(subject_n):
            idx = i*subject_n + j
            lect_hist_id = str(idx+1).zfill(10)
            lect_ctlg_id = ctlg_elms[idx][0]
            lect_DAT = now_unix_time # テストデータは20分刻みで用意する
            now_unix_time += 1200
            res.append([lect_hist_id, lect_ctlg_id, lect_DAT])

    return res

def make_atnd_info_table(pairs, hist_elms, take_n, subject_n):
    total_subject = len(hist_elms) // subject_n # 全講義の数
    res = []
    atnd_info_id = 0
    for i in range(len(pairs)):
        rndm_lect_hist_idAry = random.sample([j for j in range(total_subject)], k=take_n) # 一人当たりが履修する講義をサンプリング
        for j in range(take_n):
            student_id = pairs[i][0]
            atnd_flag = 0 # 初期値0
            lect_hist_id = rndm_lect_hist_idAry[j] * subject_n
            for k in range(subject_n):
                lect_hist_id += 1
                atnd_info_id += 1
                res.append([str(atnd_info_id).zfill(10), student_id, str(lect_hist_id).zfill(10), atnd_flag])

    return res

def make_sqlText(resDict, column_nameDict, type_nameDict):
    table_nameAry = column_nameDict.keys() # テーブル名list
    all_sentence = ""
    for table_name in table_nameAry:
        drop_sentence = f"DROP TABLE {table_name};\n"
        create_sentence = f"CREATE TABLE {table_name} ("
        flag = 0
        for column_name in column_nameDict[table_name]:
            type = type_nameDict[column_name]
            if flag == 1 and type[-11:] == "PRIMARY KEY":
                type = type[:-12]
                create_sentence += f"{column_name} {type}, "
            else:
                create_sentence += f"{column_name} {type}, "

            if type[-11:] == "PRIMARY KEY":
                flag = 1
        create_sentence = create_sentence[:-2] # 末尾のカンマ+半角スペースを除去
        create_sentence += ");\n"

        all_sentence += drop_sentence
        all_sentence += create_sentence

        for elms in resDict[table_name]:
            insert_sentence = f"INSERT INTO {table_name} VALUES ("
            for e, column_name in zip(elms, column_nameDict[table_name]):
                if type_nameDict[column_name] == "integer":
                    insert_sentence += f"{e}, "
                else:
                    insert_sentence += f"'{e}', "
            insert_sentence = insert_sentence[:-2] # 末尾のカンマ+半角スペースを除去
            insert_sentence += ");\n"
            all_sentence += insert_sentence
        all_sentence += "\n" # テーブルごとにみやすくなるように改行

    return all_sentence

def write_file(all_text, output_path):
    with open(output_path, "w") as wf:
        wf.write(all_text)

def main(data_path, subject_n, take_n, column_nameDict, type_nameDict, output_path, password_output_path):
    bar = tqdm(total=6)
    student_df = pd.read_excel(data_path, sheet_name="名前", index_col=0, header=1, usecols=[1])
    student_df.reset_index(inplace=True)
    subject_df = pd.read_excel(data_path, sheet_name="科目名", index_col=0, header=1, usecols=[1])
    subject_df.reset_index(inplace=True)
    bar.update(1)

    pairs, pairs_txt, student_elms = make_student_table(student_df)
    bar.update(1)
    ctlg_elms = make_lect_ctlg_table(subject_df, subject_n)
    bar.update(1)
    hist_elms = make_lect_hist_table(subject_df, subject_n, ctlg_elms)
    bar.update(1)
    atnd_elms = make_atnd_info_table(pairs, hist_elms, take_n, subject_n)
    bar.update(1)
    resDict = {"student": student_elms,
               "lecture_catalog": ctlg_elms,
               "lecture_history": hist_elms,
               "attendance_information": atnd_elms}
    all_text = make_sqlText(resDict, column_nameDict, type_nameDict)
    write_file(all_text, output_path)

    write_file(pairs_txt, password_output_path)
    bar.update(1)
    print("\nSuccess making data!\n")


""" 定数定義 """
DATA_PATH = os.path.join("..","data","database_testdata.xlsx")
OUTPUT_PATH = os.path.join("..","data","output.sql")
PASSWORD_INFO_PATH = os.path.join("..","data","password.txt")
SUBJECT_N = 18 # 1科目あたりの講義回数
TAKE_N = 10 # 一人あたりが履修する講義数
# TABLE_NAMES = ["student", "lecture_catalog", "lecture_history", "attendance_information"] # テーブル名
# 各テーブル名とそれに対応するカラム名
COLUMN_DICT = {"student": ["student_id", "student_name", "hashed_password"],
                "lecture_catalog": ["lecture_catalog_id", "lecture_name"],
                "lecture_history": ["lecture_history_id", "lecture_catalog_id", "lecture_date_and_time"],
                "attendance_information": ["attendance_information_id", "student_id", "lecture_history_id", "attendance_flag"]}
# カラム名に対応する型
TYPE_DICT = {"student_id": "char(10) PRIMARY KEY",
              "student_name": "varchar(30)",
              "hashed_password": "char(64)",
              "lecture_catalog_id": "varchar(100) PRIMARY KEY",
              "lecture_name": "varchar(20)",
              "lecture_history_id": "char(10) PRIMARY KEY",
              "lecture_date_and_time": "integer",
              "attendance_information_id": "char(10) PRIMARY KEY",
              "attendance_flag": "integer"}

if __name__ == "__main__":
    main(DATA_PATH, SUBJECT_N, TAKE_N, COLUMN_DICT, TYPE_DICT, OUTPUT_PATH, PASSWORD_INFO_PATH)