

### 挖掘 SimpleSection.o 的相关内容
1、gcc -c SimpleSection.c 
2、objdump -h SimpleSection.o
3、size SimpleSection.o
用来查看 ELF 文件的代码段、数据段和 BSS 段的长度

4、objdump 
-s 参数， 将所有段段内容以十六进制段方式打印出来. 
-d 参数， 所有包含指令的段的段反汇编

5、利用 nm 来查看 "SimpleSection.o" 符号
