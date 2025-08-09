import React, { useState, useEffect } from 'react';
import {useParams} from "react-router";
import axios from "axios";

interface Item {
    id: number;
    title: string;
    text: string;
}

const AnswerElement = () => {
    const { id } = useParams<{ id: string }>();
    const [item, setItem] = useState<Item | null>(null);

    useEffect(() => {
        const fetchItem = async () => {
            try {
                const response = await axios.get<Item>(
                    process.env.REACT_APP_API_URL+`/answer/get/${id}`
                );
                setItem(response.data);
            } catch (err) {
                console.error('Error fetching product:', err);
            }
        };

        fetchItem();
    }, [id]);

    return (
        <div>
            <h3>{item?.title}</h3>
            <div>{item?.text}</div>
        </div>
    )
}

export default AnswerElement
