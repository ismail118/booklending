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