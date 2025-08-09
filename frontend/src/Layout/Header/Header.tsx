import React from 'react'
import {Link} from "react-router";
import styles from "./Header.module.css"

class Header extends React.Component {
    render() {
        return (
            <div className={styles.wrap}>
                <Link to={"/"} className={styles.item}>{"Главная"}</Link>
                <Link to={"/list"} className={styles.item} >{"Список запросов"}</Link>
            </div>
        )
    }
}

export default Header
