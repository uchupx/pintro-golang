# pintro-golang

- menggunakan MYSQL/MARIDB
- jalankan sql query padahal folder sql dump dummy
- pastikan untuk membuat file config.json berdasarkan config.json.example
- sebelum build jalankan "go mod vendor" untuk mengundug asset asset yg di butuhkan


URLs
- GET /games
- POST  /games
- PUT  /games/:id
- DELETE /games/:id
- POST /users
- POST /users/sign-in
- GET /genres
- GET /publishers
- GET /platforms
- GET /regions

Noted
- Auth berlaku pada semua url untuk method POST, PUT, dan DELETE, kecuali url path "/users"
- Token tidak ada masa expired karena untuk testing


thanks for bbrumm, for shared dummy video games sql

