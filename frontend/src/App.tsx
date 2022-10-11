import React, {createContext, useContext, useEffect, useState} from 'react';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import {BrowserRouter, Route, Routes} from 'react-router-dom';
import Screen from './Views/Screen';
import APIClient from './API';
import {Edit} from './Views/Edit';
import {AddSlide} from './Views/AddSlide';
import {Layout} from './Views/Layout';
import {Slideshow} from './API/Types';
import {AlertContext, AlertData} from './Components/Alert';
import {Main} from './Views/Main';
import {Start} from './Views/Start';

export const ApiContext = createContext(new APIClient());

function App() {
    const [alertData, setAlertData] = useState<AlertData>({});
    const [selectedSlideshow, setSelectedSlideshow] = useState<Slideshow>({name: '', slides: []});
    const cli = useContext(ApiContext);

    useEffect(() => {
        cli.getSlideshows().then((res) => {
            if (res.length > 0) {
                setSelectedSlideshow(res[0]);
            }
        });
    }, []);

    return (
        <AlertContext.Provider value={{data: alertData, setData: setAlertData}}>
            <ApiContext.Provider value={new APIClient()}>
                <BrowserRouter>
                    <Routes>
                        <Route path={'/screen'} element={<Screen />} />
                        <Route path={'/'} element={<Layout />}>
                            <Route path={'/'} element={<Main />} />
                            <Route path={'/edit'} element={<Edit selectedSlideshow={selectedSlideshow} />} />
                            <Route path={'/edit/add_slide'} element={<AddSlide />} />
                            <Route path={'/start'} element={<Start />} />
                        </Route>
                    </Routes>
                </BrowserRouter>
            </ApiContext.Provider>
        </AlertContext.Provider>
    );
}

export default App;