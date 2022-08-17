import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Observable } from 'rxjs';
import { Product } from '../product';
import { ProductService } from '../product.service';

@Component({
  selector: 'app-search-product',
  templateUrl: './search-product.component.html',
  styleUrls: ['./search-product.component.css']
})
export class SearchProductComponent implements OnInit {
    searchForm: FormGroup;
    submitted=false;
    motcle:string;
    produits:Array<any>=[];
    products: Observable<Product[]>;

  
    constructor(private formBuilder:FormBuilder, private productservice:ProductService) { }
  
    ngOnInit() {
      this.reloadData();
      this.produits;
      this.searchForm=this.formBuilder.group({
        motcle:['',Validators.required],
      
      
  })
    }
    reloadData() {
      this.products = this.productservice.getProductsList();
    }
    get f(){
      return this.searchForm.controls;
  }
      onSubmit(){
        this.submitted=true;
        if(this.searchForm.invalid){
          return;
        }
     const motcle=this.searchForm.value.motcle;
     console.log(motcle);
    this.productservice.getProductByCategory(motcle).subscribe(res=>{
      this.produits=res;
      console.log(res);
    },err=>{
      console.log(err);
    }
    )
  
    }
  
  
  
    }
  