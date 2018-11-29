# Jexia code challenge: EnglishToKlingon_Converter
This project contains code to convert English word to Klingon which is a sequence of well defined hexadecimal values from the klingon alphabet set.
After displaying the corresponding hexadecimal values, the species name is fetched using the API's exposed at http://stapi.co/ 

Standard Points:
1. This repository has been made public with all commit history included;

2. To run this program, you will need first to install:

		Git
		Go

3. To run the program:
	a. Clone it in your local unix system using git.
	b. Go to main directory of project:
				$ cd startrek_englishtoklingonconverter/src/main
	c. Compile using below command:
		$ go build
	d. It will create executable binary named main
	d. For executing, run the binary along with the English name as first argument. Spaces are allowed between names.
		$ ./main Uthura
		
4. Alternative to step 3, you can also run as below:
	go run main.go Uthura
	
5. Consider each Klingon cognate letter a valid correspondence to an English letter. For example, ​D is a valid 	 correspondence of ​d
   and so on. You might notice that some letters are missing which means they are not translatable for this test purposes, then ignore the whole input;
   
6. The output contain:
	a. The translated name in Klingon written using the correspondent hexadecimal numbers according to the given table. Format:
		i. Each hexadecimal number should be separated from each other  using a single space;
		ii. If the translated name has spaces, use 0x0020 for representing each space character;
	b. The species of the given Star Trek character name using the API;
	c. The translated name and the species name separated by a new line;
	
7. To run the test cases:
	a. Execute below command from command line in each module:
		go test -v


 Output example:																									
```																
$  go run main.go Uthura									
0XF8E5 0XF8D6 0XF8E5 0XF8E1 0XF8D0
Human																

```
