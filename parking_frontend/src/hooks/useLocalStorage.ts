import {useState} from "react";


export function useLocalStorage<T>(
key:string,
initial:T
){

const [value,setValue]=useState<T>(()=>{
    
const data=localStorage.getItem(key);

return data ?
JSON.parse(data):
initial;

});


const save=(data:T)=>{

localStorage.setItem(
key,
JSON.stringify(data)
);

setValue(data);

}


return [value,save] as const;

}