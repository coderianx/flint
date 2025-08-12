# FakeData()

## FakeData() fonksiyonu, test amaçlı sahte veriler üretir.

### **Nasıl kullanılır:**
```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    flint.FakeData("data.json", 5)
    // Bu, 5 sahte kullanıcı kaydı içeren data.json adlı bir JSON dosyası oluşturur.
}
```

### **Örnek çıktı:**
```json
[
  {
    "id": 1,
    "username": "user1_GEQ",
    "email": "trijtpxb5@flintgo.dev",
    "password": "kXakAckFng",
    "hashed_password": "6f3c125a22a673d0eac1e807e39f7f2fcc7f79226023d823af94283604d006ac"
  },
  {
    "id": 2,
    "username": "user2_IzT",
    "email": "kcstgod@fakemail.com",
    "password": "Z5LVsoxKeT",
    "hashed_password": "d398e9540c590ebfb6e4196bfa461a3c79c0561888502c832e5bac3363bb58d6"
  },
  {
    "id": 3,
    "username": "user3_CHd",
    "email": "cmcktj@flintgo.dev",
    "password": "u9x0t9qc2A",
    "hashed_password": "7cc517b0447ce4db4d1bad9eb659cb2d5b49a3689d701bc733e63dfc97f7c5e5"
  },
  {
    "id": 4,
    "username": "user4_lZ3",
    "email": "dxhicxbo@flintgo.dev",
    "password": "SZ3MRwzHVK",
    "hashed_password": "ab69f50b6327bc7aca87d0614f1bab26f617966c7ff3f195b73d446de3ae558f"
  },
  {
    "id": 5,
    "username": "user5_S5i",
    "email": "emmlmlwl78@flintgo.dev",
    "password": "BZdxkpbzGn",
    "hashed_password": "4cfa0a104f256180310c0de8d6d4584d3c43602683a0f18b58d2583060f1f355"
  }
]

```

### 🧠 Açıklama:
FakeData() fonksiyonu, test amaçlı sahte veriler üretmek için kullanılır. Bu fonksiyon, belirtilen sayıda sahte kullanıcı kaydı içeren bir JSON dosyası oluşturur. Kullanıcı adı, e-posta, şifre ve şifrelenmiş şifre gibi alanlar içerir. Bu, geliştiricilerin uygulamalarını test ederken gerçek veriler kullanmadan sahte verilerle çalışmasına olanak tanır.
