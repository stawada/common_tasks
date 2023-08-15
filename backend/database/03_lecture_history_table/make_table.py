from tqdm import tqdm

# グローバル変数
TABLE_NAME = "lecture_history"
SQL_FILE_PATH = "../02_lecture_catalog_table/lecture_catalog.sql"
OUTPUT_PATH = "./lecture_history.sql"

def read_sqlFile(sql_file_path):
    lct_ctlg_idAry = []
    with open(sql_file_path, "r") as rf:
        lines = rf.readlines()
        for line in lines[3:]:
            line = line.rstrip("\n")
            els = line.split(" ")
            lecture_catalog_id = int(els[4][2:-2]) # 加工しやすいようにint型
            lct_ctlg_idAry.append(lecture_catalog_id)
    
    return lct_ctlg_idAry

def make_table(sql_file_path, table_name, output_path):
    # 2023/08/09 13:00:00のunixtimeは`1691553600`
    # これをサンプルデータのベース
    unixtime = 1691553600
    catalog_idAry = read_sqlFile(sql_file_path)
    insert_sentence = ""
    for idx, lecture_catalog_id in tqdm(enumerate(catalog_idAry)):
        # 講義履歴id
        lecture_history_id = str(idx).zfill(10)
        # 講義目録id
        lecture_catalog_id = str((lecture_catalog_id+10)//10).zfill(10)
        # 講義日時(unix time)
        # サンプルデータでは3時間おきで作る
        lecture_date_and_time = unixtime
        unixtime += 10800 # 10800s = 180m = 3h
        insert_sentence += f"INSERT INTO {table_name} VALUES ('{lecture_history_id}', '{lecture_catalog_id}', {lecture_date_and_time});\n"
    
    with open(output_path, "w") as wf:
        wf.write(f"DROP TABLE {table_name};\n")
        wf.write(f"CREATE TABLE {table_name} (lecture_history_id char(10) PRIMARY KEY, lecture_catalog_id char(10), lecture_date_and_time integer);\n\n")
        wf.write(insert_sentence)

make_table(SQL_FILE_PATH, TABLE_NAME, OUTPUT_PATH)