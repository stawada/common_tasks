import random
from tqdm import tqdm

# グローバル変数
TABLE_NAME = "lecture_catalog"
OUTPUT_PATH = "./lecture_catalog.sql"
N = 100

def make_table(n, table_name, output_path):
    insert_sentence = ""
    for idx in tqdm(range(n)):
        # 講義目録id
        lecture_catalog_id = str(idx).zfill(10)
        # 講義名
        lecture_name = "".join(random.choices([chr(i) for i in range(97,123)], k=random.randint(1,20)))
        # 講義id
        # このカラムは用途が不明なので使わない可能性あり
        lecture_id = lecture_catalog_id[::-1] # 現状はテキトーにしておく
        insert_sentence += f"INSERT INTO {table_name} VALUES ('{lecture_catalog_id}', '{lecture_name}', '{lecture_id}');\n"
    
    with open(output_path, "w") as wf:
        wf.write(f"DROP TABLE {table_name};\n")
        wf.write(f"CREATE TABLE {table_name} (lecture_catalog_id char(10) PRIMARY KEY, lecture_name varchar(20), lecture_id varchar(30));\n\n")
        wf.write(insert_sentence)

make_table(N, TABLE_NAME, OUTPUT_PATH)