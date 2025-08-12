# Register App / Kayıt Uygulaması

Bu dosya, **Flint Framework** kullanarak basit bir kullanıcı kayıt formu örneğini hem Türkçe hem İngilizce olarak açıklar.  
This file explains a simple user registration form example using **Flint Framework** in both Turkish and English.

---

## 📂 Project Structure / Proje Yapısı
```
examples/register_app/
├── main.go
├── index.html
└── README.md
```
- **main.go** → Server code / Sunucu kodu  
- **index.html** → Basic HTML registration form / Basit HTML kayıt formu  
- **README.md** → This description file / Bu açıklama dosyası  

---

## 🚀 Running / Çalıştırma

**English**
1. Install Flint framework:
```bash
go get github.com/coderianx/flint
```
2. Run the example:
```bash
go run main.go
```
3. Open in browser:
```
http://localhost:8080
````

**Türkçe**
1. Flint framework'ü yükle:
```bash
go get github.com/coderianx/flint
```
2. Örneği çalıştır:
```bash
go run main.go
```
3. Tarayıcıdan aç:
```
http://localhost:8080
```

---

## 📋 Usage / Kullanım

**English**
1. Enter a username and password in the form.
2. Click the **Submit** button.
3. The server will capture the submitted data and respond with a welcome message.

**Türkçe**
1. Açılan formda kullanıcı adı ve şifre girin.
2. **Gönder** butonuna basın.
3. Sunucu, girilen bilgileri alır ve hoş geldin mesajı döner.

---

## 📌 Notes / Notlar

**English**
- This example demonstrates how Flint’s **FormData()** method works.
- In real-world applications, passwords should be **hashed** and stored in a database.

**Türkçe**
- Bu örnek, Flint’in **FormData()** metodunu nasıl kullandığını gösterir.
- Gerçek projelerde şifreler **hashlenmeli** ve veritabanına kaydedilmelidir.
