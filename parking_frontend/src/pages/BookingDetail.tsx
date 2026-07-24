import {useEffect,useState} from "react";


export default function BookingDetail(){


const booking =
JSON.parse(
localStorage.getItem("booking")!
);


const [time,setTime]=useState("");


useEffect(()=>{


const interval=setInterval(()=>{


const start =
new Date(
booking.startTime
).getTime();


const end =
start+
booking.duration*3600000;


const diff=end-Date.now();


if(diff>0){

const h=Math.floor(
diff/3600000
);

const m=Math.floor(
(diff%3600000)/60000
);


setTime(
`Sisa ${h} jam ${m} menit`
);


}else{


setTime(
`Overtime ${
Math.floor(
Math.abs(diff)/60000
)
} menit`
);


}


},1000);



return ()=>clearInterval(interval);


},[]);



function finish(){

localStorage.removeItem(
"booking"
);

alert(
"Parkir selesai"
);

}



return (

<div className="card">

<h1>
Detail Parkir
</h1>


<p>
Slot :
{booking.slot}
</p>

<p>
Nama :
{booking.name}
</p>


<p>
Kendaraan :
{booking.vehicle}
</p>


<p>
Mulai :
{booking.startTime}
</p>


<h2>
{time}
</h2>


<button onClick={finish}>
Akhiri Parkir
</button>


</div>


)

}