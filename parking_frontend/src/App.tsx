import {
BrowserRouter,
Routes,
Route
}
from "react-router-dom";


import Home from "./pages/Home";
import BookingDetail from "./pages/BookingDetail";


function App(){

return (

<BrowserRouter>

<Routes>

<Route 
path="/"
element={<Home/>}
/>


<Route
path="/detail"
element={<BookingDetail/>}
/>


</Routes>

</BrowserRouter>

)

}

export default App;