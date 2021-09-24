# evermos-online-store

# Packages
go.mod

# Set up
1. create database
run query in file cmd\database\data\database.sql
run query in file cmd\database\data\sample_data.sql
generate tabel and insert sample data
```

2. migrate database has not ben completed 
bisa juga mamakai migrate mohon maaf strukktur migrate akan berbeda desain nya
jadi memakai run query nomor 1 saja
```
go run cmd/database/migrate.go 
go run .\cmd\database\migrate.go

go run cmd/database/migrate.go seed users
go run .\cmd\database\migrate.go seed users

go run cmd/database/migrate.go seed products
go run .\cmd\database\migrate.go seed products
```

# Features
- [X] Users as customer
- [X] Send Email When Register
- [X] User Token JWT use Redis
- [X] Products
- [X] Checkouts
- [X] Stock Minimal Carts
- [X] Stock Minimal Products
- [X] Sales Orders Status
- [ ] Cancel Orders
- [X] Migration & Seed Database
- [ ] Suppliers default use Supplier ID = 1
- [ ] Message Broker Notification for Minimum Stock


# Start 
go run main.go

# User Access 
"email": "babahmania@gmail.com",
"password": "admin123"

# Link for demo
https://158.140.191.182:50212/evermos/

http://158.140.191.182:50212/swagger/index.html

file postman colection
cmd\database\data\evermos-online.postman_collection

test and swagger end point
http://158.140.191.182:50212/swagger
https://158.140.191.182:50212/swagger

ada beberapa end point yang hanya bisa run di postman
mohon maaf setting untuk swagger nya belum bisa / selesai

note server saya berada di rumah
mohon maaf jika wifi offline sementara tidak bisa di akses 


# Solusi Acara 12.12
-Berdasarkan fakta-fakta di atas, silahkan lakukan hal-hal berikut:
1. Jelaskan menurut Anda apa yang terjadi yang menyebabkan ulasan buruk tersebut selama acara 12.12 kami dan mengapa hal itu terjadi. Letakkan ini di bagian di
your README.md file
Penyebabnya :
	a. Jumlah inventaris salah bahakan sampai negatif
		Pelanggan masih bisa checkout / order item barang tersebut.
	b. Validasi End Point checkout tidak berfungsi dengan baik
		Ketika Jumlah inventaris sudah 0 dan atau kurang dari jumlah order seharusnya tidak valid / tidak bisa checkout


2. Berdasarkan analisis Anda, ajukan solusi yang akan mencegah insiden terjadi lagi. Letakkan ini di bagian di 
your README.md file
Solusi :
	a. Validasi Database Tabel Jumlah inventaris harus >=0 dan atau bilangan positif.
		Dan menggunakan trigger untuk update Jumlah inventaris
		ketika ada insert create order update jumlah nya dan insert history inventory
	b. Perbaiki end point / api ketika checkout
		ketika jumlah pesanan lebih besar dari jumlah inventaris / stok tidak bisa checkout
		sistem penyimpanan data harus menggunakan mode transactional
		jadi ketika ada proses query yang error, semua proses simpan dibatalkan dan proses checkout pun gagal create order
	c. Di database di buat history Jumlah inventaris
		Selain untuk kebutuhan reporting keluar masuk inventaris
		ini bisa juga di jadikan refrensi untuk perhitungan jumlah stok terakhir ketika proses checkout
	d. Di buat service notif / alert (memakai message broker)
		notif bisa perhari atau hanya kirim 1 kali notif
		ini untuk meningkatkan layanan pelanggan
		- ketika stok mulai kritis / minimal
		- ketika stok checkout kurang
		- ketika membuka halaman keranjang / cart nilainya ada yang <= settingan stok keranjang, tampilkan notif kepada user
	e. Jika tidak memakai trigger ketika insert create order
		bisa dibuat stored procedure di database untuk menghitung ulang per item barang tersebut dari awal sampai akhir, data stok akhirnya akan update ke tabel inventory dengan validasi jika hasil stok akhir perhitungan adalah >=0, di jalankan di endpoint checkout ketika proses simpan data berhasil
		(mungkin ini implementasi yang kurang bagus)

3.Berdasarkan solusi yang Anda usulkan, buat Bukti Konsep yang menunjukkan secara teknis bagaimana solusi Anda akan bekerja
1. Validasi Database / Tabel Inventory
	jika di set integer positif / bilangan positif
	data minus tidak akan bisa di simpan / error query
2. Endpoint / API
	jika menggunakan mode transactional ketika update data minus akan terjadi error, dan response checkout nya menjadi gagal create order
3. History Inventaris
	akan terlihat detail data inventaris masuk/bertambah dan berkurang, stok terakhir berapa bisa di hitung
4. Update Proses perhitungan stok akhir dengan stored procedure atau trigger
	karena proses nya tidak manual, tapi hasil dari trigger/store procedure proses nya otomatis dari server database sendiri, jadi tidak manual update
	ini meminimalisasi kesalah perhitungan jumlah inventaris
	(kemungkinan salah tetep ada, ketika server database tidak bisa menjalankan proses trigger/store procedure, misalnya server nge-lock, down/mati dan atau di restart padahal sedang ada proses)
5. Service Notif / Alert Jumlah Inventaris ke Customer
    Ketika membuka halaman keranjang tampilkan notif stok yang tersedia saat ini tinggal sisa berapa
    nilainya sesuai nilai setting alert stok cart
    agar customer/user segera melakukan checkout.
6. Service Notif / Alert Jumlah Inventaris ke Supplier
	Ini bisa menambah layanan pelanggan untuk menghindari kekurangan stok ketika pelanggan order jumlahnya kurang.