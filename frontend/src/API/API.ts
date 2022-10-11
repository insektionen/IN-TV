import {Screen, Slideshow, Status} from './Types';

const baseURL = 'http://localhost:8081';

export default class APIClient {
    private async request<T>(endpoint: string, options: RequestInit): Promise<T> {
        const url = baseURL + endpoint;
        const headers = {
            'Accept': 'application/json',
        };

        const config: RequestInit = {
            ...options,
            headers: headers
        };

        return fetch(url, config)
            .then(res => {
                if (res.ok) {
                    return res.json();
                }
                throw new Error(res.status + ':' + res.statusText);
            });
    }

    public createSlideshow(sl: Slideshow) {
        return this.request<Slideshow>('/api/v1/slideshow', {
            method: 'POST',
            body: JSON.stringify(sl)
        });
    }

    public updateSlideshow(sl: Slideshow) {
        return this.request<Slideshow>('/api/v1/slideshow/' + sl.name, {
            method: 'PUT',
            body: JSON.stringify(sl)
        });
    }

    public getSlideshows() {
        return this.request<Slideshow[]>('/api/v1/slideshow', {});
    }

    public getSlideshow(name: string) {
        return this.request<Slideshow>('/api/v1/slideshow/' + name, {}).then((res) => {
            res.slides = res.slides.sort((a, b) => a.position - b.position);
            return res;
        });
    }

    public startSlideshow(slideshow: string, screens: Screen[]) {
        const data = screens.map((s) => s.name);
        return this.request('/api/v1/slideshow/' + slideshow + '/start', {
            method: 'POST',
            body: JSON.stringify(data)
        });
    }

    public stopSlideshow(slideshow: string) {
        return this.request('/api/v1/slideshow/' + slideshow + '/stop', {
            method: 'POST',
        });
    }

    public registerScreen(name: string) {
        return this.request('/api/v1/register', {
            method: 'POST',
            body: JSON.stringify({name: name})
        });
    }

    public getStatus() {
        return this.request<Status>('/api/v1/status', {}).then(res => {
            res.connected_screens.sort((a, b) => a.name.localeCompare(b.name));
            return res;
        });
    }


}
