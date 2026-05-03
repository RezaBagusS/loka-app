# LokaApp - Financial Management System 🚀

LokaApp adalah aplikasi pengelolaan keuangan modern yang dirancang untuk membantu pengguna mengalokasikan pendapatan (budgeting) dan melacak pengeluaran secara real-time dari berbagai sumber dana (Bank, E-Wallet, RDN, dsb).

## 🛠️ Tech Stack
- **Backend**: Golang (Clean Architecture)
- **Frontend**: React.js (Vite) + Tailwind CSS + shadcn/ui
- **Database**: PostgreSQL
- **State Management**: TanStack Query (React Query)
- **Infrastructure**: Docker & Docker Compose

## 📌 Fitur Utama (MVP)
- **Multi-Source Tracking**: Kelola banyak rekening/sumber dana dalam satu aplikasi.
- **Category-Based Budgeting**: Alokasikan dana ke kategori spesifik (Makan, Investasi, Cicilan, dll).
- **Data Integrity**: Menjamin akurasi saldo dengan sistem transaksi database (ACID).
- **Dashboard Visual**: Grafik ringkasan pengeluaran dan sisa budget bulanan.

## 📁 Struktur Proyek
```text
loka-app/
├── api/              # Backend (Golang)
├── web/              # Frontend (React)
├── database/         # SQL Schema & Migrations
│   └── schema.sql    # Struktur tabel awal
├── ROADMAP.md        # Rencana pengembangan jangka panjang
└── README.md         # Dokumentasi utama
```

## 🚀 Cara Memulai (Coming Soon)
1. Clone repositori:
   ```bash
   git clone https://github.com/RezaBagusS/loka-app.git
   ```
2. Jalankan infrastruktur (Docker):
   ```bash
   docker-compose up -d
   ```
3. Setup database menggunakan `database/schema.sql`.

---
*Proyek ini dikembangkan mengikuti [Roadmap Pengembangan](ROADMAP.md).*
