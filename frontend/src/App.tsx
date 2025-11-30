import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { CreateNote } from './pages/CreateNote';
import { ViewNote } from './pages/ViewNote';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<CreateNote />} />
        <Route path="/:id" element={<ViewNote />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;

