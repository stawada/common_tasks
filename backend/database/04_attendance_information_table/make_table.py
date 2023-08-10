from tqdm import tqdm

# グローバル変数
TABLE_NAME = "attendance_information"
STUDENT_SQL_PATH = "../01_student_table/student.sql"
HISTORY_SQL_PATH = "../03_lecture_history_table/lecture_history.sql"
OUTPUT_PATH = "./attendance_information.sql"

def read_sqlFile(sql_file_path):
    dataAry = []
    with open(sql_file_path, "r") as rf:
        lines = rf.readlines()
        for line in lines[3:]:
            line = line.rstrip("\n")
            els = line.split(" ")
            target_data = els[4][2:-2]
            dataAry.append(target_data)
    
    return dataAry

def make_table(std_sql_path, hst_sql_path, table_name, output_path):
    std_idAry = read_sqlFile(std_sql_path)
    hst_idAry = read_sqlFile(hst_sql_path)
    insert_sentence = ""
    for idx, hst_id in tqdm(enumerate(hst_idAry)):
        atnd_info_id = str(idx).zfill(10)
        std_id = std_idAry[idx//10]
        attend_flag = 0
        insert_sentence += f"INSERT INTO {table_name} VALUES ('{atnd_info_id}', '{std_id}', '{hst_id}', {attend_flag});\n"
    
    with open(output_path, "w") as wf:
        wf.write(f"DROP TABLE {table_name};\n")
        wf.write(f"CREATE TABLE {table_name} (attendance_information char(10) PRIMARY KEY, student_id char(10), lecture_history_id char(10), attendance_flag integer);\n\n")
        wf.write(insert_sentence)

make_table(STUDENT_SQL_PATH, HISTORY_SQL_PATH, TABLE_NAME, OUTPUT_PATH)