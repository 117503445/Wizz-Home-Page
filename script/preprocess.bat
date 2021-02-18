gf gen dao
gf pack document/create.sql,document/test.sql packed/sql.go -y
gf swagger --pack -y
gofmt -l -s -w ./