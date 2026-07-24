import {useState} from "react";
import ParkingMap from "../components/ParkingMap";
import BookingForm from "../components/BookingForm";


export default function Home(){


const [selected,setSelected]=useState<any>();


const slots=Array.from(
{length:20},
(_,i)=>({

id:i,

name:`A-${i+1}`,

occupied:i%3===0,

size:"Car"

})
);


return (

<div>


<h1>
🚗 Smart Parking System
</h1>


<ParkingMap

slots={slots}

select={setSelected}

/>



{
selected &&
<BookingForm
slot={selected}
success={()=>{
alert(
"Booking berhasil!"
)
}}
/>
}


</div>

)

}