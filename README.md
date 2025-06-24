# BOOKLENDING

App untuk pinjam buku.

## Dokumentasi API

Untuk detail lengkap mengenai API proyek ini, termasuk endpoint, parameter, dan contoh respons, silakan kunjungi dokumentasi API kami:

**[Lihat Dokumentasi API Lengkap di Postman Documenter](https://documenter.getpostman.com/view/17086351/2sB2xCiVCB)**

Atau Anda bisa langsung mengaksesnya di: `https://documenter.getpostman.com/view/17086351/2sB2xCiVCB`

---

## Setup Instruction
Run Following command :
### Install Mysql
$ make install-mysql
### Create DB In Mysql
$ make create-db
### Install golang-migrate
$ make install-migrate
### Run Migration
$ make migrate-up

### Run the Go application
$ go run main.go

---

## Table Schema
![Table (Schema Table)](bookleading.png)

---

## Project Structure
![Table (Schema Table)](Screenshot.png)

## Major Technical Decision

### PASETO token insted JWT token
PASETO is safeties and easy to use then JWT, becasue JWT give developer so much options and way to implement token that 
often lead developer that not really familiar with security to make mistake and vulnerable security issue.

### Store HashedPassword insted Naked Password
This is best practice to always store password in hashed form insted naked password for security reason

