import React, {ReactChild, useContext, useEffect, useState} from 'react';
import MQTT from 'paho-mqtt';
import {Slide} from '../API/Types';
import {ApiContext} from '../App';
import {SlideViewer} from './Slides/SlideViewer';
import {SlideThumbnail} from './Slides/SlideThumbnail';

// const SlideSVG = () => {
//     return (
//         <svg width="0">
//             <filter id="blurBackground" filterUnits="userSpaceOnUse">
//                 <feGaussianBlur stdDeviation="10" edgeMode="duplicate" />
//                 <feComponentTransfer>
//                     <feFuncA type="discrete" tableValues="1 1" />
//                 </feComponentTransfer>
//             </filter>
//         </svg>
//     );
// };

const SlideFrame = (props: { children: ReactChild | ReactChild[], isCurrent: boolean }) => {
    return <div style={{
        transition: '1s ease-out',
        width: '100%',
        height: '100%',
        opacity: props.isCurrent ? 1 : 0,
        display: props.isCurrent ? 'block' : 'none'
    }}>{props.children}</div>;
};

export const ScreenPlayer = (props: { name: string, previews?: boolean }) => {
    const [client, setClient] = useState<MQTT.Client>();
    const [slideshowName, setSlideshowName] = useState('');
    const [slides, setSlides] = useState<Slide[]>();
    const [progress, setProgress] = useState({current: 0, next: 0});
    const cli = useContext(ApiContext);

    useEffect(() => {
        if (!client) {
            connectMqtt();
        }
    }, [client]);

    useEffect(() => {
        if (slideshowName !== '' && client) {
            cli.getSlideshow(slideshowName).then((res) => {
                setSlides(res.slides);
            });
        }
    }, [slideshowName]);

    const handleMessage = (msg: MQTT.Message) => {
        const payload = JSON.parse(msg.payloadString);
        if (msg.destinationName === 'kistan/in_tv2/screen/' + props.name + '/slideshow') {
            setSlideshowName(payload.running);
        } else {
            setProgress(payload);
        }
    };

    const connectMqtt = () => {
        const cli = new MQTT.Client('server.insektionen.se', 9999, 'in-tv-client-' + props.name + Math.random().toString(16).substring(2, 8));
        cli.onConnectionLost = (err: MQTT.MQTTError) => {
            console.log('MQTT: Disconnected', err);
        };
        cli.onMessageArrived = handleMessage;

        console.log('MQTT: Connecting...');
        cli.connect({
            onSuccess: () => {
                console.log('MQTT: Connected!');
                cli.subscribe('kistan/in_tv2/screen/' + props.name + '/#');
                cli.subscribe('kistan/in_tv2/slideshow/+/change');
            },
            onFailure: (err: MQTT.ErrorWithInvocationContext) => {
                console.log('MQTT: Failed to connect:', err);
            },
            reconnect: true,
            timeout: 5
        });
        setClient(cli);
    };

    return (
        <div style={{
            overflow: 'hidden',
            cursor: 'auto',
            inset: '0px',
            height: '100%',
            width: '100%'
        }}>
            {slides ?
                slides.map(s =>
                    <SlideFrame key={s.position} isCurrent={s.position === progress.current}>
                        {props.previews ?
                            <SlideThumbnail slide={s} /> :
                            <SlideViewer slide={s}
                                isNext={s.position === progress.next}
                                isCurrent={s.position === progress.current} />}
                    </SlideFrame>)
                : null
            }
        </div>
    );
};

