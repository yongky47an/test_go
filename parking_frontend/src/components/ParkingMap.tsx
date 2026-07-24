import {
    Stage,
    Layer,
    Rect,
    Text
} from "react-konva";

import type { ParkingSlot } from "../types/parking";


interface Props {
    slots: ParkingSlot[];
    select: (slot: ParkingSlot) => void;
}


export default function ParkingMap({
    slots,
    select
}: Props) {

    return (
        <Stage width={600} height={400}>
            <Layer>

                {
                    slots.map((slot,index)=>(

                        <Rect
                            key={slot.id}
                            x={(index % 5) * 110}
                            y={Math.floor(index / 5) * 100}
                            width={90}
                            height={70}
                            fill={
                                slot.occupied
                                ? "#ef4444"
                                : "#22c55e"
                            }
                            cornerRadius={10}
                            onClick={()=>{
                                if(!slot.occupied){
                                    select(slot)
                                }
                            }}
                        />

                    ))
                }

            </Layer>
        </Stage>
    )
}