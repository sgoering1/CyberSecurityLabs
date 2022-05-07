package hscan

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
)

//==========================================================================\\

var shalookup map[string]string
var md5lookup map[string]string

func GuessSingle(sourceHash string, filename string) string {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text()

		// TODO - From the length of the hash you should know which one of these to check ...
		// add a check and logicial structure
		if len(sourceHash) == 32 {
			hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (MD5): %s\n", password)
				return password
			}
		} else if len(sourceHash) == 64 {
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (SHA-256): %s\n", password)
				return password
			}
		} else {
			log.Printf("Neither md5 or sha256")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return ""
}

func GenHashSHA(password string) {
	hash1 := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	shalookup[hash1] = password
}

func GenHashMD5(password string) {
	hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	md5lookup[hash] = password
}

func initMaps() (map[string]string, map[string]string) {
	shalookup := make(map[string]string)
	md5lookup := make(map[string]string)
	return md5lookup, shalookup
}

func GenHashMaps(filename string) (string, string) {

	//TODO
	//itterate through a file (look in the guessSingle function above)
	//rather than check for equality add each hash:passwd entry to a map SHA and MD5 where the key = hash and the value = password
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	md5lookup, shalookup = initMaps()
	for scanner.Scan() {
		password := scanner.Text()
		go GenHashSHA(password)
		go GenHashMD5(password)
	}

	log.Printf("%v", md5lookup)

	log.Printf("\n\n %v", shalookup)

	i := fmt.Sprintf("%v", md5lookup)
	k := fmt.Sprintf("%v", shalookup)
	return i, k
	//TODO at the very least use go subroutines to generate the sha and md5 hashes at the same time

	//OPTIONAL -- Can you use workers to make this even faster

	//TODO create a test in hscan_test.go so that you can time the performance of your implementation
	//Test and record the time it takes to scan to generate these Maps
	// 1. With and without using go subroutines
	// 2. Compute the time per password (hint the number of passwords for each file is listed on the site...)
}

func GenHashMap_NoGo(filename string) (string, string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	md5lookup, shalookup = initMaps()
	for scanner.Scan() {
		password := scanner.Text()
		GenHashMD5(password)
		GenHashSHA(password)
	}
	i := fmt.Sprintf("%v", md5lookup)
	k := fmt.Sprintf("%v", shalookup)
	return i, k
}

func GetSHA(hash string) (string, error) {
	password, ok := shalookup[hash]
	if ok {
		return password, nil

	} else {

		return "", errors.New("password does not exist")

	}
}

//TODO
func GetMD5(hash string) (string, error) {
	password, ok := md5lookup[hash]
	if ok {
		return password, nil
	} else {
		return "", errors.New("password does not exist")
	}
}
