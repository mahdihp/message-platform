# message-platform
سامانه پیامک


# مرحله اول ایجاد ماگریشن:

.\atlas-windows-amd64-latest.exe migrate diff InitMigration --to     "ent://ent/schema" --dev-url "postgres://admin:123456@localhost:5432/mydb?search_path=public&sslmode=disable"


# مرحله دوم بعدازاعمال تغییرات روی جدول ها و هم اکنون ایجاد ماگریشن بعدی:

.\atlas-windows-amd64-latest.exe migrate diff update_relation_01 --to "ent://ent/schema" --dev-url "docker://postgres/16/dev?search_path=public"

# مرحله سوم اعمال تغییرات:

.\atlas-windows-amd64-latest.exe  migrate apply --url "postgres://admin:123456@localhost:5432/mydb?search_path=public&sslmode=disable"

