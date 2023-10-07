# IN-TV

A system to display a slideshow on TV-Screens. Made for Kistan.

This is V2 of the whole application and it is not yet in operation for Kistan.

This application is a combined frontend and backend system. The frontend is made using React and Typescript and the
backend is made using some small frameworks in Golang.

## Configuration

Configuration is done with an YML file placed in the current working directory or `/etc/insektionen` with the
name `in_tv.yml`.

The default configuration will be written to current directory if no configuration could be loaded.

## Running

For a successful build both the frontend and backend needs to be compiled. This will create the compiled React frontend
for embedding into the Golang binary.

### 1. Frontend compilation

`cd` into the frontend directory and install all dependencies with `npm i`. Then build the resulting output
using `npm run build`.

### 2. Backend compilation

Use the go toolchain to compile the backend as normal. It will use Golangs embed feature to embed all files from the
frontend build step. This embedded code is then packaged into the final binary and served using the built-in web-server.


## Todo list
- [ ] Better documentation
   - [ ] Write more in Readme
   - [ ] Explain the folder structure
   - [ ] Explain backend, frontend
   - [ ] Explain difference about screen and computer app in frontend
   - [ ] Write documentation about the API's post, put, get methods
   - [ ] Explain what is new about IN-TV V2
   - [ ] Add idea and future work and idea list
- [ ] Implement SL info into a normal slide
- [ ] Implement SL info into a sliding banner
- [ ] Convert Backend service into background job

