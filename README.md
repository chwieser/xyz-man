# xyz-man

A small program changing the tide value of any HYPACK® XYZ-file. It simply adds a constant value to every entry in the XYZ-file.

Usage:
```

 xyz-man  input-file.xyz correction-value
        input-file.xyz      XYZ-file
        correction-value    any flot number (e.g 1.23)
       Output: input-file-(correction-value).xyz with the corrected depth values
```

eg:
```
 xyz-man.exe  input.xyz -123.3
```
generates an input---123.3.xyz file with all data entries corrected by -123.3.


This program has a rather naïve way how the input file is handled - it simply loads everything to memory before manipulation. If you run out of memory - drop me a line and I improve the program.

Enjoy,

Christian
