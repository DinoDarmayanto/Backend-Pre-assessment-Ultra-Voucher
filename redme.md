Pre Assessment Software Engineer - Back End Ultra Voucher




Terimakasih sudah tertarik menjadi bagian dari Team Software Engineer Ultra Voucher. Berikut adalah 2 problem yang kita harapkan kamu selesaikan maksimal 1x24 Jam, 
Solusi atas problem ini dapat di-submit melalui GitHub Gists dan dikirimkan tautannya kembali ke hrd@ultravoucher.co.id
 

Update terhadap result kamu akan di email maksimal 2x24 Jam oleh Ultra Voucher. 

Good Luck & Happy coding!


    1. Logic Test

Terdapat sebuah array of strings sebagai berikut:

['cook', 'save', 'taste', 'aves', 'vase', 'state', 'map']

Buatlah sebuah fungsi yang mengelompokkan sebuah array of strings di atas menjadi kumpulan anagram, dengan expected result sebagai berikut:

[
  [ 'cook' ],
  [ 'save', 'aves', 'vase' ],
  [ 'taste', 'state' ],
  [ 'map' ]
]

Catatan: dilarang menggunakan built in function seperti array.map, array.reduce, dan array.filter.

Berkiut kode yang saya berikan sesuai soal tersebut menggunakan bahasa golang.

package main

import (
	"fmt"
	"sort"
)

// fungsi untuk mengurutkan karakter pada string
func sortString(s string) string {
	characters := []rune(s)
	sort.Slice(characters, func(i, j int) bool {
		return characters[i] < characters[j]
	})
	return string(characters)
}

// fungsi untuk mengelompokkan array of strings menjadi kumpulan anagram
func groupAnagrams(strs []string) [][]string {
	anagramGroups := make(map[string][]string)

	for _, str := range strs {
		// mengurutkan karakter dalam string untuk mencari anagram
		sortedStr := sortString(str)

		// tambahkan string ke dalam map dengan kunci berupa sortedStr
		anagramGroups[sortedStr] = append(anagramGroups[sortedStr], str)
	}

	// konversi map menjadi slice dari slice untuk hasil akhir
	var result [][]string
	for _, group := range anagramGroups {
		result = append(result, group)
	}

	return result
}

func main() {
	// array of strings yang diberikan
	input := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}

	// panggil fungsi untuk mengelompokkan anagram
	result := groupAnagrams(input)

	// print hasil
	fmt.Println(result)
}

penjelasan singkat kodenya :
1.Kode di atas adalah implementasi Go untuk mengelompokkan array string menjadi kumpulan anagram. 
2.Fungsi sortString mengurutkan karakter dalam string. 
3.Fungsi groupAnagrams menggunakan map untuk mengelompokkan string berdasarkan urutan karakternya. 
4.Di loop, setiap string diurutkan dan dimasukkan ke dalam map. 
5.Hasil keluaran berupa slice dari slice string yang menunjukkan kumpulan anagram. 
6.Contoh hasil keluaran adalah [[cook] [aves save vase] [taste state] [map]].

2.Query Test 

Terdapat tabel sebagai berikut:

id          name      parent_id
1           Zaki      2
2           Ilham     NULL
3           Irwan     2
4           Arka      3

Buatlah query SQL yang menghasilkan data sebagai berikut:

id       name          parent_name
1        Zaki          Ilham
2        Ilham         NULL
3        Irwan         Ilham         
4        Arka          Irwan

Pertama-tama, saya telah membuat tabel "student" dengan 
kolom-kolom yang sesuai seperti yang dijelaskan sebelumnya di soal :

CREATE TABLE student (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    parent_id INT
);

Kemudian, saya telah memasukkan data ke dalam tabel "student" sesuai dengan 
data yang diberikan di soal :

INSERT INTO student (id, name, parent_id) VALUES (1, 'Zaki', 2);
INSERT INTO student (id, name, parent_id) VALUES (2, 'Ilham', NULL);
INSERT INTO student (id, name, parent_id) VALUES (3, 'Irwan', 2);
INSERT INTO student (id, name, parent_id) VALUES (4, 'Arka', 3);

Setelah data dimasukkan, saya ingin menghasilkan output yang menampilkan Id, nama 
siswa, dan nama orang tua siswa (jika ada) dengan menggunakan query JOIN :

SELECT 
    t1.id, 
    t1.name,
    t2.name AS parent_name
FROM 
    student t1
LEFT JOIN 
    student t2 ON t1.parent_id = t2.id;

Penjelasan query:

1.SELECT t1.id, t1.name: Memilih kolom id dan name dari tabel utama "student" dengan alias
 t1.
2.t2.name AS parent_name: Menggabungkan tabel utama "student" dengan tabel dirinya sendiri 
dengan alias t2 berdasarkan kolom parent_id pada t1 dan kolom id pada t2. Lalu memilih kolom
 name dari tabel t2 dengan alias parent_name.
3.FROM student t1: Menentukan tabel utama dengan alias t1.
4.LEFT JOIN student t2 ON t1.parent_id = t2.id: Melakukan LEFT JOIN antara tabel utama "student" 
dengan tabel dirinya sendiri dengan alias t2 berdasarkan kolom parent_id pada t1 dan kolom id pada t2.

Dengan query di atas, saya dapat mendapatkan hasil sebagai berikut:



id  | name  | parent_name
----|-------|------------
1   | Zaki  | Ilham
2   | Ilham | NULL
3   | Irwan | Ilham
4   | Arka  | Irwan




“Becoming the Leader Market in Voucher Industry”