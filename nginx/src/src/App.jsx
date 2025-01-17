import { BrowserRouter, Routes, Route, Link, NavLink } from 'react-router-dom';
import SearchPage from './pages/SearchPage'
import ResultPage from './pages/ResultPage'

const App = () => {
	return (
		<BrowserRouter>
			<Routes>
				<Route path="/" element={<SearchPage />} />
				<Route path="/result" element={<ResultPage />} />
			</Routes>
		</BrowserRouter>
	)
}

export default App