// 連接到 admin 資料庫以進行驗證
db = db.getSiblingDB('admin');

// 使用 root 使用者和密碼進行驗證，從環境變數取得 MONGO_USERNAME 和 MONGO_PASSWORD
db.auth(process.env.MONGO_INITDB_ROOT_USERNAME, process.env.MONGO_INITDB_ROOT_PASSWORD);

// 切換到 cp_tracker 資料庫
db = db.getSiblingDB('cp_tracker');

// 建立一些初始集合
db.createCollection('users');

// 建立用戶集合的索引
db.users.createIndex({ "email": 1 }, { unique: true });
