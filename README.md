# Hacktiv8 Golang Final Project 4 Team 8

## Link Deploy API

## Link Postman Documentation

## Anggota

1. Mochamad Suhri Ainur Rifky (GLNG-KS04-001)
2. Raden Muhammad Yudie Sanjaya (GLNG-KS04-016) :x:
3. Varrel Marcellius (GLNG-KS04-021)

## Pembagian Tugas

1. Mochamad Suhri Ainur Rifky (GLNG-KS04-001)

- API Product
- API Transaction History
- Postman (Collection, Environment, Documentation)

2. Raden Muhammad Yudie Sanjaya (GLNG-KS04-016) :x:


3. Varrel Marcellius (GLNG-KS04-021)

- Setup project
- API User
- API Category
- Readme
- Deploy
- Unit Test

# Cara Install

1. run `docker compose up` untuk menjalankan database
2. run `go run main.go seeder` untuk menjalankan aplikasi dan seeder admin
3. run `go run main.go` untuk menjalankan aplikasi jika sudah pernah seeder admin

## Credential Admin
```
email: admin@gmail.com
password: admin123
```

# List Route

## Users

- `POST` - `/users/register`. Digunakan customer untuk registrasi
- `POST` - `/users/login`. Digunakan admin dan customer untuk login
- `PATCH` - `/users/topup`. Digunakan admin dan customer untuk topup balance

## Categories

- `POST` - `/categories`. Digunakan admin untuk membuat kategori baru
- `GET` - `/categories`. Digunakan admin untuk mendapat semua data kategori
- `PATCH` - `/categories/:categoryId`. Digunakan admin untuk mengupdate kategori
- `DELETE` - `/categories/:categoryId`. Digunakan admin untuk menhapus kategori

## Products

- `POST` - `/products`. Digunakan admin untuk membuat produk baru
- `GET` - `/products`. Digunakan admin dan pengguna untuk mendapat semua data produk
- `PUT` - `/products/:productId`. Digunakan admin untuk mengupdate produk
- `DELETE` - `/products/:productId`. Digunakan admin untuk menghapus produk

## Transaction Histories

- `POST` - `/transactions`. Digunakan customer untuk melakukan transaksi
- `GET` - `/transactions/my-transactions`. Digunakan customer untuk mendapat semua data transaksi
- `GET` - `/transactions/user-transactions`. Digunakan admin untuk melihat seluruh data transaksi customer