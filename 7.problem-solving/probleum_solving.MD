# Problem Solving

## Question ❓
Anda memiliki sebuah API yang mengembalikan data JSON. Bagaimana cara Anda 
mengambil data dari API tersebut menggunakan Golang?

## Answer 💡
Untuk mengambil data dari API, kita akan membuat sebuah API Call berbentuk HTTP Request dengan mengingat bahwa expected response yang dikirim akan berbentuk JSON.
Di Golang, untuk melakukan HTTP Request sebenarnya cukup mudah, kita cukup menggunakan package `net/http`, package standard yang sudah disediakan Golang. Untuk methodnya, kita bisa memilih NewRequest atau NewRequestWithContext. Perbedaannya hanya ada di args functionnya, dimana NewRequestWithContext memerlukan args tambahan yaitu context yang berguna untuk mengirim signal/informasi ke dalam requestnya, misalnya signal untuk cancellation dan timeout.
Pastikan juga kita sudah mendefine sebuah `struct` yang memiliki struktur yang sama dengan response JSON yang diharapkan, agar proses `decoding` datanya bisa berjalan dengan baik.

## Question ❓
Anda mendapat laporan bahwa ada masalah dengan performa aplikasi yang 
dibangun dengan Golang. Bagaimana cara Anda menemukan dan memperbaiki 
masalah tersebut?

## Answer 💡
Adapun langkah-langkah yang bisa diterapkan sebagai berikut:
1. Mengecek source code, atau dengan kata lain memahami kiranya mengapa aplikasi bisa berjalan dengan lambat. Apakah ada masalah di aplikasinya atau ada penyebab dari external, misalnya resource server yang sudah termakan habis.
2. Melakukan cek kinerja aplikasi dengan `profiling` dan `benchmarking`. Kedua hal ini bisa dicoba, mengingat Golang sudah mensupport kedua hal ini. Dengan melakukan profiling dan benchmarking, kita bisa mengetahui berapa waktu yang diperlukan untuk mengeksekusi semua code kita.
3. Menerapkan `caching`. Caching diperlukan agar jumlah request yang menuju ke database berkurang. Hal ini tentunya bisa mempercepat aplikasi kita. Namun perlu diingat, tidak semua data bisa di-caching, mengingat dalam caching ada waktu jeda untuk me-refresh datanya. Jadi, untuk menjaga data consistency, pastikan hanya meng-caching data yang jarang berubah.
4. Menggunakan `Goroutines` di aplikasi. Goroutines adalah salah satu fitur bawaan Go untuk menghandle concurrency, dimana nantinya tiap task/function yang menerapkan Goroutines akan dijalankan di mini thread yang berbeda-beda, sehingga code nya akan semakin cepat dieksekusi.
5. Setelah melakukan upaya-upaya di atas, baiknya kita kembali mengecek performa aplikasi melalui log maupun melakukan load test di API yang kita bangun.  

## Question ❓
Anda diminta untuk membangun sebuah aplikasi backend dengan Golang. Aplikasi 
tersebut harus dapat menerima request HTTP dan mengembalikan response berupa 
data JSON. Bagaimana cara Anda memulai dan mengembangkan aplikasi tersebut? 
Apa saja yang perlu diperhatikan dalam proses pengembangan aplikasi tersebut?

## Answer 💡
Berikut cara minimal untuk membangun sebuah API di Golang:
1. Mendifinisikan route dan handler. Route berperan sebagai url endpoint yang akan diakses client, sementara handler berfungsi untuk mengolah request dan response yang akan diterima balik oleh client.
2. Membuat http Server. Membuat route dan server tidaklah cukup, setelah itu kita wajib membuat server dengan port yang kita tentukan, agar request yang dikirim oleh client bisa kita terima.
3. Memproses request dan mengembalikan response. Kita bisa mengembalikan response JSON sesuai dengan yang kita mau. Langkah 1-3 ini bisa dilakukan hanya dengan standard library dari Go, yaitu net/http. Meskipun ada banyak sekali library lain yang bisa digunakan, seperti Gin, Echo dan Fiber.
4. Membuat unit test. Unit test juga sangat penting dibuat untuk mengetahui API yang kita bangun sudah mampu mengembalikan response seperti yang kita minta.
5. Melakukan integration test. Membuat unit test tidaklah cukup, setelah itu kita juga wajib mengetest API secara keseluruhan menggunakan tools seperti Postman maupun melalui curl di terminal.
6. Setelah API dipastikan mampu berjalan, langkah selanjutnya adalah memantau API dan melakukan penjagaan terhadap API yang telah dibangun.
