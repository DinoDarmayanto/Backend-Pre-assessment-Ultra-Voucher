-- Pertama-tama, saya telah membuat tabel "student" dengan 
-- kolom-kolom yang sesuai seperti yang dijelaskan sebelumnya di soal :

CREATE TABLE student (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    parent_id INT
);

-- Kemudian, saya telah memasukkan data ke dalam tabel "student" sesuai dengan 
-- data yang diberikan di soal :

INSERT INTO student (id, name, parent_id) VALUES (1, 'Zaki', 2);
INSERT INTO student (id, name, parent_id) VALUES (2, 'Ilham', NULL);
INSERT INTO student (id, name, parent_id) VALUES (3, 'Irwan', 2);
INSERT INTO student (id, name, parent_id) VALUES (4, 'Arka', 3);

-- Setelah data dimasukkan, saya ingin menghasilkan output yang menampilkan Id, nama 
-- siswa, dan nama orang tua siswa (jika ada) dengan menggunakan query JOIN :

SELECT 
    t1.id, 
    t1.name,
    t2.name AS parent_name
FROM 
    student t1
LEFT JOIN 
    student t2 ON t1.parent_id = t2.id;























