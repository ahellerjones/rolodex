import sqlite3
class SQLiteHandler: 
    def __init__(self, db_file):
        self.db_file = db_file
        self.connection = None

    def connect(self):
        self.connection = sqlite3.connect(self.db_file)

    def disconnect(self):
        if self.connection:
            self.connection.close()
            self.connection = None

    def execute_query(self, query):
        cursor = self.connection.cursor()
        cursor.execute(query)
        result = cursor.fetchall()
        cursor.close()
        return result


