create /app
create /app/port
create /app/port/user ":8082"
create /app/port/company ":8081"
create /app/port/transaction ":8080"
create /database
create /database/uri "mongodb+srv://CashbagMe:Cashbag@cluster0.epe8y.gcp.mongodb.net/Cluster0?retryWrites=true&w=majority"
create /database/name
create /database/name/user "CashBagUser"
create /database/name/transaction "CashBagTransaction"
create /database/name/company "CashBagCompany"
create /database/test
create /database/test/user "CashBagUserTest"
create /database/test/transaction "CashBagTransactionTest"
create /database/test/company "CashBagCompanyTest"
create /redis
create /redis/uri "redis:6379"
create /redis/pass ""
create /grpc
create /grpc/address
create /grpc/address/user "localhost"
create /grpc/address/company "localhost"
create /grpc/address/transaction "localhost"
create /grpc/port
create /grpc/port/user ":9002"
create /grpc/port/company ":9001"
create /grpc/port/transaction ":9000"
