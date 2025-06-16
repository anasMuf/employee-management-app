# Employee API Documentation

REST API untuk manajemen data employee yang mendukung operasi Create, Read, Update, dan Delete (CRUD).

## Informasi API

- **Versi**: 1.0.0
- **Base URL**: `https://employee-management-app-a71ae14c9e4a.herokuapp.com/`
- **Kontak**: Anas Mufti (anas.muhammadakbar@gmail.com)

## Endpoints

### 1. GET /employee
Mengambil seluruh data employee dari database.

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "name": "Andi",
    "email": "andi@example.com",
    "phone": "081010101010"
  }
]
```

### 2. POST /employee
Menambahkan employee baru ke database.

**Request Body:**
```json
{
  "name": "Budi",
  "email": "budi@example.com",
  "phone": "087777777777"
}
```

**Response (201 Created):**
```json
{
  "id": 2,
  "name": "Budi",
  "email": "budi@example.com",
  "phone": "087777777777"
}
```

**Response Error (400 Bad Request):**
Data input tidak valid.

### 3. GET /employee/{id}
Mengambil data employee berdasarkan ID tertentu.

**Parameter:**
- `id` (integer, required): ID employee

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "Andi",
  "email": "andi@example.com",
  "phone": "081010101010"
}
```

**Response Error (404 Not Found):**
Employee tidak ditemukan.

### 4. PUT /employee/{id}
Mengupdate data employee berdasarkan ID.

**Parameter:**
- `id` (integer, required): ID employee

**Request Body:**
```json
{
  "name": "Andi Updated",
  "email": "andi.updated@example.com",
  "phone": "081010101011"
}
```

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "Andi Updated",
  "email": "andi.updated@example.com",
  "phone": "081010101011"
}
```

**Response Error:**
- `400 Bad Request`: Data input tidak valid
- `404 Not Found`: Employee tidak ditemukan

### 5. DELETE /employee/{id}
Menghapus data employee berdasarkan ID.

**Parameter:**
- `id` (integer, required): ID employee

**Response (204 No Content):**
Employee berhasil dihapus.

**Response Error (404 Not Found):**
Employee tidak ditemukan.

## Schema Data

### Employee (Response)
```json
{
  "id": integer,
  "name": string,
  "email": string,
  "phone": string
}
```

### EmployeeInput (Request Body)
```json
{
  "name": string,
  "email": string,
  "phone": string
}
```

**Catatan:** Seluruh field pada EmployeeInput bersifat required (wajib diisi).

## Contoh Penggunaan

### Menggunakan cURL

**Ambil semua employee:**
```bash
curl -X GET https://employee-management-app-a71ae14c9e4a.herokuapp.com/employee
```

**Tambah employee baru:**
```bash
curl -X POST https://employee-management-app-a71ae14c9e4a.herokuapp.com/employee \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","phone":"081234567890"}'
```

**Update employee:**
```bash
curl -X PUT https://employee-management-app-a71ae14c9e4a.herokuapp.com/employee/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"John Updated","email":"john.updated@example.com","phone":"081234567891"}'
```

**Hapus employee:**
```bash
curl -X DELETE https://employee-management-app-a71ae14c9e4a.herokuapp.com/employee/1
```

## Status Codes

- `200 OK`: Request berhasil
- `201 Created`: Resource berhasil dibuat
- `204 No Content`: Resource berhasil dihapus
- `400 Bad Request`: Data input tidak valid
- `404 Not Found`: Resource tidak ditemukan