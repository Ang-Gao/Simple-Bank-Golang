//@c
/*************************************************************************
 *@description:
 *@author: Jason Gao
 *@date: 2023-03-02 22:57:02
*************************************************************************/
//define variable
let var_text : string = "I am testing!";
let num : number = 0;
let sentence : string = `We can use template to print content: ${var_text}`;

//array
let list1 : number[] = [1, 2, 3, 4];
let list2 : Array<number> = [1,2,3,4];

//tuple
let person1 : [string, number] = ['name',0];

//enum
enum Color {RED = 'RED COLOR',GREEN ='',YELLOW = ''};
let c : Color = Color.RED;
console.log(c); //不初始化枚举的值，默认值就是index

//也可以像js一样赋予弱类型（soft type,放弃类型检查）(any)可以赋予这个var任意类型的值/func
let randomValue : any = 10;
randomValue = function randomValue(){}

//unknown type, any放弃类型检查，可以使用任何方法，但是unknown必须“类型收缩”(通过as 某 type)之后才能使用其响应的方法
let my_var : unknown = 10;

//多类型
let multiType : string | number;
multiType = 'hi';
multiType = 1;


//function, n2?:number, 加了个问号就代表这个参数可以是optional的，调用的时候可以不写这个参
//最后的那个number指的是返回值的类型
//默认undefined如果没赋值的话，并且undefined也算是boolean的子集(undefined是所有类型的子集)
//所以如果是undefined的话，如果用if，则能n2代表了false
//以及，可以给parameter默认值，在func直接声明值
function add(n1 : number = 10, n2 ?: number) : number{
  //n2 exists
  if(n2)
    return n1 + n2;
  else
    return n1;
}
console.log(add(1, 1));

//interface
interface Person{
  firstName : string;
  lastName : string;
}

function fullName(p : Person){
  console.log(`${p.firstName} ${p.lastName}`); //反引号才能让template生效
}
let person = {
  firstName : 'Jason',
  lastName : 'Gao'
}
fullName(person);

//class(OOP) and access modifiers
//access modifiers(public,private,protected)
/*
public --> free accessibility
ts默认就是public权限，不像java，默认是在同一包里才能访问
private --> only accessed within a class
protected --> accessibility within a class and the classes derived from it(在同一类里，或者是这个类的子类)
*/
class Students{
  firstName : string;
  lastName : string;
  constructor(firstName : string, lastName : string){
    this.firstName = firstName;
    this.lastName = lastName;
  }
  getFirstName() : string { return this.firstName; }
}
let stu1 = new Students('Angx',"Gao");
console.log(stu1.getFirstName());

class Employee implements Person{
  firstName: string;
  lastName: string;
  constructor(firstName : string, lastName : string){
    this.firstName = firstName;
    this.lastName = lastName;
  }
}

class Manage extends Employee{
  private manager : boolean;
  constructor(manager:boolean,firstName:string,lastName:string){
    super(firstName,lastName);
    this.manager = manager;
  }
  getManager(){
    console.log(`This is manager? ${this.manager}, ${this.firstName} ${this.lastName}`);
  }
}
let m1 = new Manage(true,"J","Gao");
m1.getManager();
