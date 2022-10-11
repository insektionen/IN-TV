import {Alert as BSAlert} from 'react-bootstrap';
import {Variant} from 'react-bootstrap/types';
import {createContext, useContext, useEffect, useState} from 'react';

export const AlertContext = createContext<AlertContextContent>({
    data: {}, setData: () => {
        console.log('Using default Context value for AlertContext');
    }
});

interface AlertContextContent {
    data: AlertData,
    setData: (data: AlertData) => void
}

export interface AlertData {
    text?: string,
    variant?: Variant
    timeout?: number
}

export const Alert = () => {
    const [showing, setShowing] = useState(false);
    const {data: alertData} = useContext(AlertContext);

    useEffect(() => {
        setShowing(true);
    }, [alertData]);

    if (!alertData.text || alertData.text === '' || !showing) {
        return null;
    }

    if (alertData.timeout) {
        setTimeout(() => setShowing(false), alertData.timeout * 1000);
    }

    return (
        <BSAlert variant={alertData.variant ?? 'secondary'}>{alertData.text ?? ''}</BSAlert>
    );
};