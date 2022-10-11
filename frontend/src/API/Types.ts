export interface Slide {
    type: string
    timeout: number
    position: number
    data: any
}

export interface Slideshow {
    name: string
    slides: Slide[]
}

export interface Screen {
    name: string
    last_seen: number
}

export interface Status {
    running_slideshows: string[]
    connected_screens: Screen[]
}