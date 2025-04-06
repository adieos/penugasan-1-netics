| Nama     | NRP |
|:--------:|:------:|
| Athallah Rajendra Wibisono    |   5025231170  |

# Teknis Pengerjaan

Implementasikan modul CI/CD ini pada sebuah sistem server sederhana, dengan detail sebagai berikut
1. Buatlah API publik dengan endpoint /health yang menampilkan informasi sebagai berikut:
```
{
  "nama": "Athallah Rajendra Wibisono",
  "nrp": "5025231170",
  "status": "UP",
  “timestamp”: time // Current time
  "uptime": time    // Server uptime
}
```
Bahasa pemrograman dan teknologi yang digunakan dibebaskan kepada peserta.
3. Lakukan deployment API tersebut dalam bentuk container (Docker Multi-stage) pada VPS publik.
4. Lakukan proses CI/CD menggunakan GitHub Actions untuk melakukan otomasi proses deployment API. Terapkan juga best practices untuk menjaga kualitas environment CI/CD
5. Dokumentasikan pengerjaan di sebuah laporan berbentuk Markdown pada repositori peserta masing-masing.

# Langkah Pengerjaan

1. Pertama, kita buat dulu API-nya. Disini, saya menggunakan Go sebagai bahasa pemrogramannya, dibantu dengan Gin sebagai frameworknya.

![image](https://github.com/user-attachments/assets/81616116-17a0-42f6-b07b-a36329841461)

endpoint `/health` akan menampilkan data, sedangkan endpoint lainnya (alias 404) akan menampilkan pesan error. Selain itu, saya juga melakukan simulasi apabila app menggunakan variabel dari env:

![image](https://github.com/user-attachments/assets/a514fc52-138e-4022-a804-dd46178b0d5e)

![image](https://github.com/user-attachments/assets/29a05044-1428-470a-9113-1b0ab0e720ae)

2. Setelah itu, kita buat Dockerfile dan docker-compose nya

![image](https://github.com/user-attachments/assets/6283d514-0cdd-4b11-a4a7-d016118cbd5f)

Di dalam Dockerfile, saya menggunakan `golang:alpine` sebagai base image karena lebih ringan, sehingga akan lebih cepat untuk di-build. Kemudian, binary di-build dan port 8888 di-expose ke luar image.

