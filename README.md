# REST API Bioskop

> Version 1.0.0

OpenAPI untuk API pengelolaan data bioskop.

- Server: `<host>/api` (Server lokal / Server Publik)


## Path Table

| Method | Path | Description |
| --- | --- | --- |
| GET | [/](#get) | Cek status API |
| GET | [/bioskop](#getbioskop) | Mendapatkan daftar bioskop |
| POST | [/bioskop](#postbioskop) | Menambahkan bioskop |
| GET | [/bioskop/:id](#getbioskopid) | Mendapatkan detail bioskop |
| PUT | [/bioskop/:id](#putbioskopid) | Memperbarui bioskop |
| DELETE | [/bioskop/:id](#deletebioskopid) | Menghapus bioskop |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |
| BioskopId | [#/components/parameters/BioskopId](#componentsparametersbioskopid) | ID bioskop |
| Bioskop | [#/components/schemas/Bioskop](#componentsschemasbioskop) |  |
| BioskopRequest | [#/components/schemas/BioskopRequest](#componentsschemasbioskoprequest) |  |
| ApiResponse | [#/components/schemas/ApiResponse](#componentsschemasapiresponse) |  |
| BioskopResponse | [#/components/schemas/BioskopResponse](#componentsschemasbioskopresponse) |  |
| BioskopListResponse | [#/components/schemas/BioskopListResponse](#componentsschemasbioskoplistresponse) |  |
| DeleteResponse | [#/components/schemas/DeleteResponse](#componentsschemasdeleteresponse) |  |
| HomeResponse | [#/components/schemas/HomeResponse](#componentsschemashomeresponse) |  |
| ErrorResponse | [#/components/schemas/ErrorResponse](#componentsschemaserrorresponse) |  |
| ValidationError | [#/components/responses/ValidationError](#componentsresponsesvalidationerror) | Input tidak valid, termasuk nama/lokasi kosong atau rating di luar rentang 0 sampai 5. |
| ServerError | [#/components/responses/ServerError](#componentsresponsesservererror) | Kesalahan database atau data tidak ditemukan. Implementasi saat ini mengembalikan status 500 untuk kondisi data tidak ditemukan. |

## Path Details

***

### [GET]/

- Summary
  Cek status API

- Operation id
  getApiHome

#### Responses

- 200 API berjalan

`application/json`

```typescript
{
  message: string
}
```

- Example

```json
{
  "message": "API Home berjalan"
}
```

***

### [GET]/bioskop

- Summary
  Mendapatkan daftar bioskop

- Operation id
  getBioskopList

#### Responses

- 200 Daftar bioskop berhasil diperoleh

`application/json`

```typescript
{
  message: string
  data: {
    id: integer
    nama: string
    lokasi: string
    rating: number //format: float //minimum: 0 //maximum: 5
  }[]
}
```

- Example

```json
{
  "message": "Berhasil mendapatkan data",
  "data": [
    {
      "id": 1,
      "nama": "Bioskop Contoh",
      "lokasi": "Bandung",
      "rating": 4.5
    }
  ]
}
```

- 500 Kesalahan database atau data tidak ditemukan. Implementasi saat ini mengembalikan status 500 untuk kondisi data tidak ditemukan.

`application/json`

```typescript
{
  message: string
  data: string | null
}
```

- Example

```json
{
  "message": "Gagal mengambil data bioskop",
  "data": "Data tidak tersedia"
}
```

***

### [POST]/bioskop

- Summary
  Menambahkan bioskop

- Operation id
  createBioskop

#### RequestBody

*Required*

- application/json

```typescript
{
  // Nama bioskop; tidak boleh kosong atau hanya spasi.
  nama: string //minLength: 1
  // Lokasi bioskop; tidak boleh kosong atau hanya spasi.
  lokasi: string //minLength: 1
  rating?: number //format: float //minimum: 0 //maximum: 5 //default: 0
}
```

- Example

```json
{
  "nama": "Bioskop Contoh",
  "lokasi": "Bandung",
  "rating": 4.5
}
```

#### Responses

- 201 Bioskop berhasil ditambahkan

`application/json`

```typescript
{
  message: string
  data: {
    id: integer
    nama: string
    lokasi: string
    rating: number //format: float //minimum: 0 //maximum: 5
  }
}
```

- Example

```json
{
  "message": "Berhasil menambahkan bioskop",
  "data": {
    "id": 1,
    "nama": "Bioskop Contoh",
    "lokasi": "Bandung",
    "rating": 4.5
  }
}
```

- 400 Input tidak valid, termasuk nama/lokasi kosong atau rating di luar rentang 0 sampai 5.

`application/json`

```typescript
{
  message: string
  data: string | null
}
```

- Examples

  - requiredFields

```json
{
  "value": {
    "message": "Nama dan lokasi wajib diisi",
    "data": null
  }
}
```

  - ratingRange

```json
{
  "value": {
    "message": "Rating harus antara 0 sampai 5",
    "data": null
  }
}
```

- 500 Kesalahan database atau data tidak ditemukan. Implementasi saat ini mengembalikan status 500 untuk kondisi data tidak ditemukan.

`application/json`

```typescript
{
  message: string
  data: string | null
}
```

- Example

```json
{
  "message": "Gagal mengambil data bioskop",
  "data": "Data tidak tersedia"
}
```

***

### [GET]/bioskop/:id

- Summary
  Mendapatkan detail bioskop

- Operation id
  getBioskopById

#### Parameters(Path)

```typescript
// ID bioskop
id: integer //minimum: 1
```

#### Responses

- 200 Detail bioskop berhasil diperoleh

`application/json`

```typescript
{
  message: string
  data: {
    id: integer
    nama: string
    lokasi: string
    rating: number //format: float //minimum: 0 //maximum: 5
  }
}
```

- Example

```json
{
  "message": "Berhasil mendapatkan data",
  "data": {
    "id": 1,
    "nama": "Bioskop Contoh",
    "lokasi": "Bandung",
    "rating": 4.5
  }
}
```

- 500 Kesalahan database atau data tidak ditemukan. Implementasi saat ini mengembalikan status 500 untuk kondisi data tidak ditemukan.

`application/json`

```typescript
{
  message: string
  data: string | null
}
```

- Example

```json
{
  "message": "Gagal mengambil data bioskop",
  "data": "Data tidak tersedia"
}
```

***

### [PUT]/bioskop/:id

- Summary
  Memperbarui bioskop

- Operation id
  updateBioskop

#### Parameters(Path)

```typescript
// ID bioskop
id: integer //minimum: 1
```

#### RequestBody

*Required*

- application/json

```typescript
{
  // Nama bioskop; tidak boleh kosong atau hanya spasi.
  nama: string //minLength: 1
  // Lokasi bioskop; tidak boleh kosong atau hanya spasi.
  lokasi: string //minLength: 1
  rating?: number //format: float //minimum: 0 //maximum: 5 //default: 0
}
```

- Example

```json
{
  "nama": "Bioskop Contoh Baru",
  "lokasi": "Jakarta",
  "rating": 4
}
```

#### Responses

- 201 Bioskop berhasil diperbarui (status aktual dari handler)

`application/json`

```typescript
{
  message: string
  data: {
    id: integer
    nama: string
    lokasi: string
    rating: number //format: float //minimum: 0 //maximum: 5
  }
}
```

- Example

```json
{
  "message": "Berhasil melakukan perubahan bioskop",
  "data": {
    "id": 1,
    "nama": "Bioskop Contoh Baru",
    "lokasi": "Jakarta",
    "rating": 4
  }
}
```

- 400 Input tidak valid, termasuk nama/lokasi kosong atau rating di luar rentang 0 sampai 5.

`application/json`

```typescript
{
  message: string
  data: string | null
}
```

- Examples

  - requiredFields

```json
{
  "value": {
    "message": "Nama dan lokasi wajib diisi",
    "data": null
  }
}
```

  - ratingRange

```json
{
  "value": {
    "message": "Rating harus antara 0 sampai 5",
    "data": null
  }
}
```

- 500 Kesalahan database atau data tidak ditemukan. Implementasi saat ini mengembalikan status 500 untuk kondisi data tidak ditemukan.

`application/json`

```typescript
{
  message: string
  data: string | null
}
```

- Example

```json
{
  "message": "Gagal mengambil data bioskop",
  "data": "Data tidak tersedia"
}
```

***

### [DELETE]/bioskop/:id

- Summary
  Menghapus bioskop

- Operation id
  deleteBioskop

#### Parameters(Path)

```typescript
// ID bioskop
id: integer //minimum: 1
```

#### Responses

- 201 Bioskop berhasil dihapus (status aktual dari handler)

`application/json`

```typescript
{
  message: string
  data: null
}
```

- Example

```json
{
  "message": "Berhasil menghapus bioskop",
  "data": null
}
```

- 500 Kesalahan database atau data tidak ditemukan. Implementasi saat ini mengembalikan status 500 untuk kondisi data tidak ditemukan.

`application/json`

```typescript
{
  message: string
  data: string | null
}
```

- Example

```json
{
  "message": "Gagal mengambil data bioskop",
  "data": "Data tidak tersedia"
}
```

## References

### #/components/parameters/BioskopId

```typescript
// ID bioskop
id: integer //minimum: 1
```

### #/components/schemas/Bioskop

```typescript
{
  id: integer
  nama: string
  lokasi: string
  rating: number //format: float //minimum: 0 //maximum: 5
}
```

### #/components/schemas/BioskopRequest

```typescript
{
  // Nama bioskop; tidak boleh kosong atau hanya spasi.
  nama: string //minLength: 1
  // Lokasi bioskop; tidak boleh kosong atau hanya spasi.
  lokasi: string //minLength: 1
  rating?: number //format: float //minimum: 0 //maximum: 5 //default: 0
}
```

### #/components/schemas/ApiResponse

```typescript
{
  message: string
  data: any
}
```

### #/components/schemas/BioskopResponse

```typescript
{
  message: string
  data: {
    id: integer
    nama: string
    lokasi: string
    rating: number //format: float //minimum: 0 //maximum: 5
  }
}
```

### #/components/schemas/BioskopListResponse

```typescript
{
  message: string
  data: {
    id: integer
    nama: string
    lokasi: string
    rating: number //format: float //minimum: 0 //maximum: 5
  }[]
}
```

### #/components/schemas/DeleteResponse

```typescript
{
  message: string
  data: null
}
```

### #/components/schemas/HomeResponse

```typescript
{
  message: string
}
```

### #/components/schemas/ErrorResponse

```typescript
{
  message: string
  data: string | null
}
```

### #/components/responses/ValidationError

- application/json

```typescript
{
  message: string
  data: string | null
}
```

- Examples

  - requiredFields

```json
{
  "value": {
    "message": "Nama dan lokasi wajib diisi",
    "data": null
  }
}
```

  - ratingRange

```json
{
  "value": {
    "message": "Rating harus antara 0 sampai 5",
    "data": null
  }
}
```

### #/components/responses/ServerError

- application/json

```typescript
{
  message: string
  data: string | null
}
```

- Example

```json
{
  "message": "Gagal mengambil data bioskop",
  "data": "Data tidak tersedia"
}
```
