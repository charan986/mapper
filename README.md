# Mapper package helps capitalizes *only* every third alphanumeric character of a string, 
 
## Below shows the usage of mapper package
 
 ```sh
 func main() {
   s := NewSkipString(3, "Aspiration.com")
   mapper.MapString(&s)
   fmt.Println(s)
}
```


  ## Notes 
  if you use
	Aspiration.com as input as shown, 
	the function returns
	asPirAtiOn.cOm
  
  Please note: 
- Characters other than each third should be downcased
- For the purposes of counting and capitalizing every three characters, consider only alphanumeric
	  characters, ie [a-z][A-Z][0-9]. 
