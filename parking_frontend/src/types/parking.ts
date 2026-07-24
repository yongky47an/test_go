export interface ParkingSlot {
    id: number;
    name: string;
    occupied: boolean;
    size: string;
}


export interface Booking {
    slot: string;
    name: string;
    vehicle: string;
    duration: number;
    startTime: string;
}