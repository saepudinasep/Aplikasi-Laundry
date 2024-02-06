----------------------- Customer -----------------------------------------------
-- Select Customer
SELECT * FROM mst_customers;

-- Select Customer Berdasarkan Id
SELECT id_customer, name, no_telp, alamat FROM mst_customers WHERE id_customer = '1';

-- Insert Customer
INSERT INTO mst_customers (id_customer, name, no_telp, alamat)
VALUES (1, 'Jessica', '0812654987', 'Bandung'),
(2, 'Michael', '0821345678', 'Jakarta'),
(3, 'Emily', '0854321098', 'Surabaya'),
(4, 'Daniel', '0813987654', 'Yogyakarta'),
(5, 'Sophia', '0897654321', 'Semarang');

-- Update Customer
UPDATE mst_customers SET name = 'Jessica', no_telp = '0812654987', alamat = 'Jakarta'
WHERE id_customer = '1';

-- Delete Customer
DELETE FROM mst_customers WHERE id_customer = '1';

--------------------------------- Layanan ---------------------------------------

-- Select Semua Layanan
SELECT * FROM mst_layanan;

-- Select Layanan Berdasarkan Id
SELECT id_layanan, nama_layanan, harga, satuan FROM mst_layanan WHERE id_layanan = '5';

-- Insert Layanan
INSERT INTO mst_layanan (id_layanan, nama_layanan, harga, satuan)
VALUES (1, 'Cuci + Setrika', 7000, 'KG'),
(2, 'Laundry Bedcover', 50000, 'Buah'),
(3, 'Laundry Boneka', 25000, 'Buah'),
(4, 'Laundry Jas', 35000, 'Pcs'),
(5, 'Laundry Sepatu', 30000, 'Pasang');

-- Update Layanan
UPDATE mst_layanan SET nama_layanan = 'Laundry Sepatu', harga = 25000, satuan = 'Pasang'
WHERE id_layanan = '5';

-- Delete Layanan
DELETE FROM mst_layanan WHERE id_layanan = '5';


--------------------------- Order ---------------------------------------------------

-- Select Semua Order
SELECT id_order, customer_id, tanggal_masuk, tanggal_keluar, penerima FROM trx_order;

-- Select Order Berdasarkan Id
SELECT id_order, customer_id, tanggal_masuk, tanggal_keluar, penerima FROM trx_order
WHERE id_order = '1';

-- Insert Order
INSERT INTO trx_order (customer_id, tanggal_masuk, tanggal_keluar, penerima)
VALUES (1, '18-08-2022', '20-08-2022', 'Mirna') RETURNING id_order;


--------------------------- Order Detail ----------------------------------------------

-- Select Semua Order Detail
SELECT id_order_detail, order_id, layanan_id, quantity FROM trx_order_detail;

-- Select Order Detail Berdasarkan Id
SELECT id_order_detail, order_id, layanan_id, quantity FROM trx_order_detail
WHERE order_id = '1';

-- Insert Order Detail
INSERT INTO trx_order_detail (order_id, layanan_id, quantity)
VALUES (1, 1, 5) RETURNING id_order_detail;

INSERT INTO trx_order_detail (order_id, layanan_id, quantity)
VALUES (1, 2, 1) RETURNING id_order_detail;

INSERT INTO trx_order_detail (order_id, layanan_id, quantity)
VALUES (1, 3, 2) RETURNING id_order_detail;