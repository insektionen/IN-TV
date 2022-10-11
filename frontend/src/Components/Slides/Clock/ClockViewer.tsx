import {Slide} from '../../../API/Types';
import {useEffect, useState} from 'react';

export const ClockViewer = (props: { slide: Slide }) => {
    const [currentTime, setCurrentTime] = useState('');
    const [currentDate, setCurrentDate] = useState('');
    const [timer, setTimer] = useState<NodeJS.Timer>();

    useEffect(() => {
        if (!timer) {
            updateTime();
            const t = setInterval(updateTime, 1000);
            setTimer(t);
        }
    }, [timer]);

    const zeroTime = (i: number) => {
        let res = i.toString();
        if (i < 10) {
            // add zero in front of numbers < 10
            res = '0' + i.toString();
        }
        return res;
    };

    const updateTime = () => {
        const now = new Date();
        const year = now.getFullYear();
        // Javascript Date returns month from 0-11
        const month = zeroTime(now.getMonth() + 1);
        const day = zeroTime(now.getDate());
        setCurrentDate(year + '-' + month + '-' + day);

        const h = now.getHours();
        const m = zeroTime(now.getMinutes());
        const s = zeroTime(now.getSeconds());
        let t = h + ':' + m;
        if (props.slide.data.with_seconds ?? false) {
            t = t + ':' + s;
        }
        setCurrentTime(t);
    };

    return (
        <div style={{
            height: '100%',
            color: 'white',
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center'
        }}>
            <div>
                {props.slide.data.with_date ?
                    <h2 style={{
                        fontSize: '900%',
                        width: '100%',
                        textAlign: 'center'
                    }}>{currentDate}</h2> : null}
                <h1 style={{
                    fontSize: '2500%',
                    color: 'white',
                    width: '100%',
                    textAlign: 'center',
                }}>{currentTime}</h1>
            </div>
        </div>
    );
};