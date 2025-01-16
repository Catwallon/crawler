import { BrowserRouter, Routes, Route, Link, NavLink } from 'react-router-dom';
import SearchPage from './pages/SearchPage'
import ResultPage from './pages/ResultPage'

function App() {
	return (
		<BrowserRouter>
			<Routes>
				<Route path="/" element={<SearchPage />} />
				<Route path="/result" element={<ResultPage />} />
				<Route path="/result:query?" element={<ResultPage />} />
			</Routes>
		</BrowserRouter>
	)
}

export default App