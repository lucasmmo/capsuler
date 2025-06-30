import {
  createBrowserRouter,
  RouterProvider,
} from 'react-router'

import './App.css'

const router = createBrowserRouter([
  {
    path: "/",
  }
])

function App() {

  return (
    <RouterProvider router={router} />
  )
}

export default App
