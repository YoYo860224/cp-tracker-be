// 連接到 admin 資料庫以進行驗證
db = db.getSiblingDB('admin');

// 使用 root 使用者和密碼進行驗證，從環境變數取得 MONGO_USERNAME 和 MONGO_PASSWORD
db.auth(process.env.MONGO_INITDB_ROOT_USERNAME, process.env.MONGO_INITDB_ROOT_PASSWORD);

// 切換到 cp_tracker 資料庫
db = db.getSiblingDB('cp_tracker');

// 建立用戶集合
if (!db.getCollectionNames().includes('user')) {
    db.createCollection('user');
}
db.users.createIndex({ "email": 1 }, { unique: true });

// 建立邀請碼集合
if (!db.getCollectionNames().includes('invite_code')) {
    db.createCollection('invite_code');
}
db.invite_codes.createIndex({ "code": 1 }, { unique: true });

// 建立 user_data 集合
if (!db.getCollectionNames().includes('user_data')) {
    db.createCollection('user_data');
}
db.user_data.createIndex({ "uid": 1 }, { unique: true });
