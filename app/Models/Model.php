<?php

namespace App\Models;

use App\Traits\DefaultDatetimeFormat;
use Illuminate\Database\Eloquent\Model as EloquentModel;

class Model extends EloquentModel
{
    use DefaultDatetimeFormat;
}
