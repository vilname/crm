import React, { useState, useEffect } from 'react';
import {Link} from "react-router";

interface Item {
    id: number;
    title: string;
    text: string;
}

interface Pagination {
    page: number;
    limit: number;
    total: number;
}

const AnswerList = () => {
    const [items, setItems] = useState([])
    const [pagination, setPagination] = useState<Pagination>({
        'page': 0,
        'limit': 0,
        'total': 0,
    })

    const [currentPage, setCurrentPage] = useState(1);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch(
                    process.env.REACT_APP_API_URL+`/answer/list?page=${currentPage}&limit=${pagination.limit}`
                )

                const result = await response.json()
                setItems(result.data);
                setPagination(result.pagination)
                // setTotalItems(data.total);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        }

        fetchData();
    }, [currentPage, pagination.limit])

    const totalPages = Math.ceil(pagination.total / pagination.limit);
    const checkPaginate = (pageNumber: number) => setCurrentPage(pageNumber);;

    return (
        <div className="pagination-container">
            <h1>Список элементов с пагинацией</h1>

            {/* Список элементов */}
            <ul className="items-list">
                {items.map((item: Item) => (
                    <div>
                        <h2><Link to={`/get/${item.id}`}>{item.title}</Link></h2>
                        <div>{item.text}</div>
                    </div>
                ))}
            </ul>

            {/* Пагинация */}
            <div className="pagination">
                <button
                    onClick={() => checkPaginate(pagination.page - 1)}
                    disabled={pagination.page === 1}
                >
                    Назад
                </button>

                {Array.from({ length: totalPages }, (_, i) => i + 1).map((number) => (
                    <button
                        key={number}
                        onClick={() => checkPaginate(number)}
                        className={pagination.page === number ? 'active' : ''}
                    >
                        {number}
                    </button>
                ))}

                <button
                    onClick={() => checkPaginate(pagination.page + 1)}
                    disabled={pagination.page === totalPages}
                >
                    Вперед
                </button>
            </div>

            <div className="page-info">
                Страница {pagination.page} из {totalPages} | Всего элементов: {pagination.total}
            </div>
        </div>
    )
}

export default AnswerList