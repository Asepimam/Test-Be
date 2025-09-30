# PracticalTest API Response Example

Berikut adalah contoh hasil response API yang akan dihasilkan untuk PracticalTest:

```json
{
    "Meta": {
        "count": 15
    },
    "data": [
        {
            "office_name": "UID JATIM",
            "room_name": "Ruang Fatmawati",
            "capacity": 55,
            "usage_percent": 60.68,
            "consume_nominal": 2350000,
            "type_consume": [
                {"name": "Snack Siang", "count": 55, "amount": 1100000},
                {"name": "Makan Siang", "count": 25, "amount": 750000},
                {"name": "Snack Sore", "count": 25, "amount": 500000}
            ]
        },
        // ...data lainnya sesuai contoh di atas...
    ],
    "message": "Summary data retrieved successfully",
    "status": "success"
}
```

## Penjelasan Struktur
- **Meta.count**: Jumlah total data.
- **data**: List data ruangan, berisi nama kantor, nama ruangan, kapasitas, persentase penggunaan, nominal konsumsi, dan detail konsumsi per jenis.
- **type_consume**: Rincian konsumsi (nama, jumlah, nominal).
- **message**: Pesan status response.
- **status**: Status response (success/gagal).

Contoh di atas dapat digunakan sebagai referensi hasil API untuk pengujian atau dokumentasi PracticalTest.

## Contoh Request

Untuk mendapatkan data summary, gunakan perintah berikut:

```bash
curl -X GET http://localhost:8080/summary
```