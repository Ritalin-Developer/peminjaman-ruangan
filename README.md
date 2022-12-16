# Peminjaman Ruangan App
![License](https://img.shields.io/badge/license-closed-green)
![App version](https://img.shields.io/badge/version--app-v1.0.0-brightgreen)
![Go version](https://img.shields.io/badge/Go-v1.19-lightblue)
![Docker version](https://img.shields.io/badge/docker--version-20.10.21-blue)

Merupakan aplikasi yang ditujukan untuk peminjaman ruangan kampus berbasis web, meskipun pada penerapannya kali ini kami tidak berfokus pada tampilan antar muka user. Kami ingin mendevelop API Service yang nantinya ekstensible, dapat digunakan oleh pihak ketiga, memiliki dokumentasi yang mumpuni, dan bahkan API ini dapat diimplementasikan menjadi model berbayar. Jadi hanya pihak ketiga yang memiliki token yang dapat mengakses API Gateway service kami.

## Containerization

Layanan kami menerapkan konsep Containerization, yang artinya Backend, Database, dan Frontend berjalan dalam suatu wadah yang dapat di deploy di server mana saja tanpa menghadapi dilemma "Will this work on my machine?". Beberapa poin keuntungan lainnya adalah:

- Consistency

Memanfaatkan platform yang bekerja dengan cara yang sama di berbagai lingkungan menghilangkan begitu banyak stres. Seluruh tim kami bekerja dengan cara yang sama, terlepas dari server, mesin, atau sistem operasi yang mereka gunakan.

- Automation

Ada begitu banyak tugas yang, sebagai developer, bisa menjadi repetitif dan monoton jika dilakukan secara manual. Kontainer Docker memungkinkan Anda untuk menjadwalkan berbagai tugas agar terjadi saat dibutuhkan, tanpa intervensi manual dari manusia. Ini menghemat waktu, tenaga, dan meringankan beban kerja developer.

- Stability

Docker didasarkan pada Linux dan, dengan demikian, memiliki kernel Linux di setiap wadah, terlepas dari sistem yang dijalankannya. Di masa lalu, ini mungkin menyebabkan beberapa masalah stabilitas kecil saat menjalankan kontainer di sistem Mac atau Windows. Hari-hari ini, meskipun Docker sering memperbarui, lingkungan tetap stabil di sistem atau perangkat apa pun. Tidak perlu tiba-tiba kembali ke pembaruan sebelumnya atau panik karena masalah kompatibilitas yang tidak terduga.

- Save Space

Pendahulu wadah adalah Mesin Virtual (VM). VM bekerja dengan cara yang mirip dengan wadah, tetapi mengambil server fisik dan meludahkannya ke lingkungan virtual, menggunakan ruang server fisik dalam jumlah besar dan banyak memori. Kontainer Docker hanya menggunakan kode untuk aplikasi dan dependensinya dan dapat berjalan sepenuhnya di Cloud, yang berarti mereka jauh lebih kecil dan meniadakan persyaratan untuk server fisik yang besar.

## Project Structure

Berikut merupakan struktur dari project Peminjaman Ruangan

1. Database

Menggunakan PostgreSQL database service, ditujukan agar dapat kedepannya lebih scalable, selain itu PostgreSQL memiliki keamanan yang baik, database yang bersifat Open Source, dan extensible.

2. Backend

Menggunakan Golang sebagai backend service. Bahasa pemrograman ini sangat cocok untuk menangani server, backend and micro service. Bahasa ini bersifat static type, sangat ringan dan cepat.

3. Frontend

Menggunakan basic HTML, CSS, dan JS untuk web interface. 

## Deployment Step

1. Install docker, tutorial lengkap dapat mengunjungi [link ini](https://docs.docker.com/engine/install/)  
2. Clone repository ini
3. Konfigurasi `backend/app.env.example` sesuai dengan value yang diinginkan, lalu rename menjadi `app.env`

```sh
# Sample 
PORT=1118
ENVIRONMENT=local
VERSION=v1.0.0
APP_NAME=peminjaman-ruangan
POSTGRES_HOST=localhost // gunakan 'database' jika ingin mendeploy service, 'localhost' merupakan database service yang sudah di expose
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=samplepassword
POSTGRES_DATABASE=peminjaman-ruangan
GORM_LOG=true
POSTGRES_MIN_CONN=1
POSTGRES_MAX_CONN=2
ALLOWERD_ORIGIN=*
SECRET_KEY=t0ps3cr3t
TOKEN_LIFETIME_MIN=3600
```

4. Jalankan syntax 

```sh
cd backend/migration/bin
cp ../../app.env . && ./backend-migration
docker compose build && docker compose up -d && docker compose ps -a
```

syntax ini akan nge-build service yang ada lalu menjalankannya di background dan nge-list status service tersebut.

5. Jika ingin melihat log service, jalankan syntax 

```sh
docker logs -f servicename
```

Ubah `servicename` menjadi `backend`, `frontend`, ataupun `database`

## Contributor

- 1922021 - Muhamad Arie
- 1922004 - Pritami Sergio

### References

- https://blog.iron.io/docker-containers-the-pros-and-cons-of-docker/