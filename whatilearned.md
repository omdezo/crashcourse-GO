Ignore this file (it's just for me to note tings down to go back to it if needed)

what is go and why it's matter: 
 compiled - typed - object oriented by it's own way 
 it's fast since it's doesn't have (TL) like the ASP.NET  
 perfect for general use and microservices 
 flixable 

I have started on my first code (FirstCode.go) test things out. 
I have start also use Make commands i found it really intersting i can see it really helpful in real projects for set commands. 

Course 1: 
    printing methods .. what i have done here is just testing the basic opretion methods and printing.

course 2: 
    Variables, Strings & Numbers .. what i have done here is just declare the variables and using it 
    My observation : there is somthings unique here i have tried to declare a value and not using it and here is the suprise it's giving me an error instead of just warning wich making me in setuation no dead code. 
    otherwise everything close to C#.

course 3: 
    Printing & Formatting Strings.. Intersting topic to discuss about the thing is I expect that it will be most likly as pyton way [f"{name} {age}"]
    but instead it's as C copy usnig [% specifiers]. because the go langauge is trendy langauge so it's little bit we should ask why they didn't use "Interpolation". the first thing in my mind goes like "wait - so why google Still Doesn't Have String Interpolation for GO langauge? " . So, as 2026 programmer ofcourse I have asked my boy "Claude" about it and he said there is actual Proposals but got Declined. also he said "The primary reasons are the existing fmt.Sprintf function already provides robust string formatting capabilities and language complexity" and so on. and then claude mention also about wat rob pike said "Once you understand the building blocks, you shouldn't need help from the reference manual to write interesting code." . Well, i think i got my answer.

course 4:
    Arrays & Slices .. it's just about the ways to declare and initialize an array also how to slice it in two ways "slice using array method  & range method" 
    "arrays.go" 

course 5: 
    It was about the standard library. basically Go ships with a lot out of the box, i touched strings and sort. with strings i used the usual stuff - Contains, ReplaceAll, ToUpper, Index, Split. the interesting thing i noticed is that none of these change the original value, strings are immutable in Go, everything returns a new value. i confirmed this myself by printing the original string after all the operations and it was unchanged.

    for sort i sorted ints and strings with sort.Ints and sort.Strings, then used SearchInts and SearchStrings to find indexes. one thing to be careful about here - the Search functions only work correctly if the slice is already sorted first. if it's not you don't get an error, you just get a wrong index back silently.

course 6: 
    It was about arrays and slices. arrays are fixed size and the size is actually part of the type itself, so [3]string and [5]string are completely different types in Go, not the same thing with different lengths. slices are the flexible version, they use arrays under the hood but you don't deal with the size directly. append doesn't modify the slice in place, it returns a new one. for range slicing the end index is always exclusive so [1:3] gives you index 1 and 2 only, not 3.

after the 6th cousrse i don't think i need type everyting now at least i have the vision how this langauge it's works. But i will write down here if there is somthing i didn't expect.

there is somthin intersting also about multiple_return_values it's fix problem error handling that i have to know if there is problem in the code in during compiling which means it's not like c / c# / c++ /ts /java / js .. so if there is an error it will force me to handle it not leting me compile. 