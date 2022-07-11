mysql  -u"$MYSQL_USER" -p"$MYSQL_PWD" db < ../../db/migrations/0_init.up.sql
mysql  -u"$MYSQL_USER" -p"$MYSQL_PWD" db < ../../db/migrations/0_create_books.up.sql