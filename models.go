package main

func getAllShipmentsFromDB() []*Shipment {
	rows, err := DB.Query("SELECT * FROM Shipments")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var allShipments []*Shipment
	for rows.Next() {
		s := new(Shipment)
		err := rows.Scan(&s.ID, &s.SenderName, &s.SenderEmail, &s.SenderAddress, &s.SenderCountryCode, &s.RecipientName, &s.RecipientEmail, &s.RecipientAddress, &s.RecipientCountryCode, &s.Weight, &s.Price)
		if err != nil {
			return nil
		}
		allShipments = append(allShipments, s)
	}

	if err := rows.Err(); err != nil {
		return nil
	}
	return allShipments
}

func getShipmentFromDB(id string) Shipment {

	result, err := DB.Query("SELECT * FROM `Shipments` WHERE id = ?", id)
	if err != nil {
		panic(err)
	}

	defer result.Close()

	var shipment Shipment

	for result.Next() {
		err := result.Scan(&shipment.ID, &shipment.SenderName, &shipment.SenderEmail, &shipment.SenderAddress, &shipment.SenderCountryCode, &shipment.RecipientName, &shipment.RecipientEmail, &shipment.RecipientAddress, &shipment.RecipientCountryCode, &shipment.Weight, &shipment.Price)
		if err != nil {
			panic(err)
		}
	}

	return shipment
}

func insertDataInDB(shipment Shipment) {
	insert, err := DB.Query("INSERT INTO Shipments(SenderName, SenderEmail, SenderAddress, SenderCountryCode, RecipientName, RecipientEmail, RecipientAddress, RecipientCountryCode, Weight, Price) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", shipment.SenderName, shipment.SenderEmail, shipment.SenderAddress, shipment.SenderCountryCode, shipment.RecipientName, shipment.RecipientEmail, shipment.RecipientAddress, shipment.RecipientCountryCode, shipment.Weight, shipment.Price)
	if err != nil {
		panic(err)
	}

	defer insert.Close()
}