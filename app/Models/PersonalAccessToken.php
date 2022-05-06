<?php

namespace App\Models;

use Laravel\Sanctum\PersonalAccessToken as Model;

class PersonalAccessToken extends Model
{
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'token', 'abilities', 'expired_at'
    ];

    /**
     * The attributes that should be cast to native types.
     *
     * @var array
     */
    protected $casts = [
        'abilities' => 'json',
        'last_used_at' => 'datetime',
        'expired_at' => 'datetime',
    ];
}
