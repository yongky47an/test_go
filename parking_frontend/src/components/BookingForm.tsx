import {useState} from "react";
import {ParkingSlot,Booking} from "../types/parking";


interface Props{

slot:ParkingSlot;

success:(data:Booking)=>void;

}


export default function BookingForm({
slot,
success
}:Props){


const [form,setForm]=useState({

name:"",
vehicle:"",
duration:1

});


function submit(){

const booking:Booking={

slot:slot.name,

name:form.name,

vehicle:form.vehicle,

duration:Number(form.duration),

startTime:new Date().toISOString()

};


localStorage.setItem(
"booking",
JSON.stringify(booking)
);


success(booking);

}



return (

<div className="card">


<h2>
Booking {slot.name}
</h2>


<input
placeholder="Nama"
onChange={
e=>setForm({
...form,
name:e.target.value
})
}
/>


<input
placeholder="Nomor Kendaraan"
onChange={
e=>setForm({
...form,
vehicle:e.target.value
})
}
/>


<input
type="number"
placeholder="Durasi Jam"
onChange={
e=>setForm({
...form,
duration:Number(e.target.value)
})
}
/>


<button onClick={submit}>
Pesan Parkir
</button>


</div>

)

}