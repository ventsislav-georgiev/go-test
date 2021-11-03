# Benchmarks

## Contents:
* C# (.NET v6) vs Go - [Exceptions](#exceptions)
* C# (.NET v6) vs Go - [JSON](#json)
* Go - [Method Invocation](#method-invocation)

<br/>

# Exceptions:
* C# code:
<img width="715" alt="image" src="https://user-images.githubusercontent.com/5616486/144600962-f494d407-bf25-4a4e-b6ac-201d55f33d96.jpg">

* C# results:
<img width="535" alt="image" src="https://user-images.githubusercontent.com/5616486/144601092-605f866b-8557-4526-94d8-9f8ddfcd92bc.jpg">

* Go code:
<img width="535" alt="image" src="https://user-images.githubusercontent.com/5616486/144602580-5e837f1e-1d97-4298-b434-7bc7522a6062.png">

* Go results:
<img width="793" alt="image" src="https://user-images.githubusercontent.com/5616486/144603002-9f24a0b0-d05b-4500-9454-37fefd5973fe.png">

## Summary:
* C# 10ns vs 11us - Try not to use exceptions in the hot paths of your code.
* Go 4ns vs 150ns - Panic in Go is generally for non user defined errors. All others are simply values passed via the `error` interface.

<br/>
<br/>

# JSON:
* C# code:
<img width="715" alt="image" src="https://user-images.githubusercontent.com/5616486/144603377-c5a233a5-92dd-49c8-adea-94b56f932ba4.jpg">

* C# results:
<img width="535" alt="image" src="https://user-images.githubusercontent.com/5616486/144603518-b941599c-18a3-4246-8541-ae15d8d2163a.jpg">

* Go code:
<img width="535" alt="image" src="https://user-images.githubusercontent.com/5616486/144603583-68277817-faa9-47f2-a79d-ac082b6d2f67.jpg">

* Go results:
<img width="842" alt="image" src="https://user-images.githubusercontent.com/5616486/144603802-c768d731-1365-4935-ba78-422756eef547.png">

## Summary:
The new JSON serializer in .NET 6 with code generation is about the same speed as the default in Go. Easyjson also uses the code generation aproach and is 4 times faster.

<br/>
<br/>

# Method Invocation:
* Go code:
<img width="489" alt="image" src="https://user-images.githubusercontent.com/5616486/144610896-b9ddcfc6-1995-49bf-9f22-bcd0639c15c0.png">
<img width="451" alt="image" src="https://user-images.githubusercontent.com/5616486/144610838-2950a756-5b63-479c-897e-c165186e428c.png">

* Go results:
<img width="691" alt="image" src="https://user-images.githubusercontent.com/5616486/144610766-1f359b82-9c4a-4cc2-97e7-38076f6b81fd.png">

## Summary:
The method invocation on type vs interface vs empty interface with type switch or type assertion have little to no difference in performance. The fastest seems to be to pass a struct as interface. Maybe the method table is smaller in this case 🤷‍♂️.

<br/>
<br/>
