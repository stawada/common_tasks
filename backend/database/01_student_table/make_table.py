import hashlib
import random
from tqdm import tqdm

# グローバル変数
STUDENT_TABLE_NAME = "student"
PASSWORD_LIST = "./password_list.txt"
STUDENT_TABLE_PATH = "./student.sql"
TABLE_NAME = "student"
N = 10 # 講義履歴テーブルの1/10を想定


# 学生テーブル
def make_student_table():
    insert_students = ""
    passwordAry = ""
    for idx in tqdm(range(N)):
        # 学生番号
        up_char = random.choice([chr(i) for i in range(97,123)]).upper() # ランダムに1文字のアルファベットを取得
        student_id = up_char + str(idx).zfill(9)
        
        # 学生の名前
        student_name = "".join(random.choices([chr(i) for i in range(97,123)], k=random.randint(1,30))) # 適当に名前を生成
        
        # パスワード生成(ハッシュ化済み)
        password = "".join(random.choices([chr(i) for i in range(97,123)], k=10))
        hashed_pw = hashlib.sha256(password.encode()).hexdigest()
        
        insert_students += f"INSERT INTO {TABLE_NAME} values ('{student_id}', '{student_name}', '{hashed_pw}');\n"
        passwordAry += student_id + " " + password + "\n"
        
    # テキストファイルに追加書き込み
    with open(STUDENT_TABLE_PATH, "w") as wf:
        wf.write(f"DROP TABLE {TABLE_NAME};\n")
        wf.write(f"CREATE TABLE {TABLE_NAME} (student_id varchar(10) PRIMARY KEY, student_name varchar(50), hashed_password char(64));\n\n")
        wf.write(insert_students)
    print(f"Make {STUDENT_TABLE_PATH}!")
    with open(PASSWORD_LIST, "w") as wf:
        wf.write(passwordAry)
    print(f"Make {PASSWORD_LIST}!")

make_student_table()